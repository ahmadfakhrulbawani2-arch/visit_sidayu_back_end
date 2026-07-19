package controllers

import (
	"context"
	"net/http"
	"os"
	cfg "simple_go_gin_gorm_postgres_be_template/internal/config"
	res "simple_go_gin_gorm_postgres_be_template/internal/helpers"
	"simple_go_gin_gorm_postgres_be_template/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"
)

func initImageKit() *imagekit.Client {
	client := imagekit.NewClient(
		option.WithPrivateKey(os.Getenv("IMAGEKIT_PRIVATE_KEY")),
	)
	return &client
}

func UploadImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		res.RespError(ctx, http.StatusBadRequest, "Image file must be provided", err)
		return
	}

	defer file.Close()

	fileName := header.Filename
	ik := initImageKit()
	uploadRes, err := ik.Files.Upload(context.TODO(), imagekit.FileUploadParams{
		File:     file,
		FileName: fileName,
	})

	if err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Failed to uplaod image to Imagekit, this may happenned because not enough storage, bad connection, etc.", err)
	}

	newImage := models.Image{
		FileID:   uploadRes.FileID,
		Name:     fileName,
		ImageUrl: uploadRes.URL,
	}

	err = cfg.DB.Create(&newImage).Error
	if err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Failed to save image metadata to database, this may happened because not enough database storage, bad database connection, etc.", err)
		return
	}

	res.RespSuccess(ctx, http.StatusCreated, "Image uploaded!", newImage, "", nil)
}
