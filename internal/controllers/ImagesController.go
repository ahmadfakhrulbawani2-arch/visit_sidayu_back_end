package controllers

import (
	"context"
	"errors"
	"net/http"
	"os"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/models"

	cfg "visit-sidayu-backend/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"gorm.io/gorm"
)

func initImageKit() *imagekit.Client {
	client := imagekit.NewClient(
		option.WithPrivateKey(os.Getenv("IMAGEKIT_PRIVATE_KEY")),
	)
	return &client
}

// POST /api/v1/images (form/app)
func UploadImage(ctx *gin.Context) {
	// check authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Image file must be provided", err)
		return
	}

	defer file.Close()

	filename := header.Filename
	ik := initImageKit()
	uploadRes, err := ik.Files.Upload(context.TODO(), imagekit.FileUploadParams{
		File:     file,
		FileName: filename,
	})

	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to uplaod image to Imagekit, this may happenned because not enough storage, bad connection, etc.", err)
	}
	customName := ctx.DefaultPostForm("custom_name", filename)
	newImage := models.CreateImages{
		Name:       filename,
		ImageUrl:   uploadRes.URL,
		FileID:     uploadRes.FileID,
		CustomName: customName,
	}

	err = cfg.DB.Create(&newImage).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to save image metadata to database, this may happened because not enough database storage, bad database connection, etc.", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusCreated, "Image uploaded!", newImage, "", nil)
}

// DELETE /api/v1/images/:id (json)
func DeleteImage(ctx *gin.Context) {
	// check authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	var image models.Images
	paramsId := ctx.Param("id")
	parsedParamsId, err := uuid.Parse(paramsId)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Can not parse parameters", err)
		return
	}

	err = cfg.DB.First(&image, "id = ?", parsedParamsId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Image not found", nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch image", err)
		return
	}

	imagekitFileID := image.FileID
	ik := initImageKit()
	err = ik.Files.Delete(context.TODO(), imagekitFileID)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to uplaod image to Imagekit, this may happenned due to imagekit server error, bad connection, etc.", err)
		return
	}

	err = cfg.DB.Unscoped().Delete(&image).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Unexpected server error to delete image", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Successfully delete image!", image, "", nil)
}

// PUT /api/v1/images/:id (form/app), ensuring complete data we use PUT instead of PATCH (actually to differentiate json and formfile payload)
func UpdateImage(ctx *gin.Context) {
	// check authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	var image models.Images
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Image file must be provided", err)
		return
	}

	defer file.Close()

	filename := header.Filename
	ik := initImageKit()
	updateRes, err := ik.Files.Upload(context.TODO(), imagekit.FileUploadParams{
		File:     file,
		FileName: filename,
	})

	paramsId := ctx.Param("id")
	parsedParamsId, err := uuid.Parse(paramsId)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Can not parse parameters", err)
		return
	}

	var errNotFound = cfg.DB.First(&image, "id = ?", parsedParamsId)
	if errNotFound != nil {
		hp.RespError(ctx, http.StatusNotFound, "Can not found image data", nil, "Data not found in database")
		return
	}

	customName := ctx.DefaultPostForm("custom_name", filename)
	updatedImage := models.CreateImages{
		Name:       filename,
		ImageUrl:   updateRes.URL,
		FileID:     updateRes.FileID,
		CustomName: customName,
	}

	err = cfg.DB.Model(&image).Updates(updatedImage).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to update data image into database, this can happened due to not enough database capacity, database bad connection, etc.", err)
		return
	}

	err = cfg.DB.First(&image, "id = ?", parsedParamsId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Image not found", nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch image", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Successfully updated image!", image, "", nil)

}

// GET /api/v1/images/:id (json)
func GetImageById(ctx *gin.Context) {
	var image models.Images
	paramsId := ctx.Param("id")
	parsedParamsId, err := uuid.Parse(paramsId)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Can not parse parameters", err)
		return
	}

	err = cfg.DB.First(&image, "id = ?", parsedParamsId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Image not found", nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch image", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Data image fetched successfully", image, "", nil)
}

// GET /api/v1/images?search=john&limit=10&page=1
func GetImages(ctx *gin.Context) {
	var images []models.Images

	query := cfg.DB.Model(&models.Images{})
	search := ctx.Query("search")
	if search != "" {
		// only search prefix for lightweight db load
		query = query.Where("customname ILIKE ? OR name ILIKE ? OR description ILIKE ?", search+"%", search+"%", search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)

	err := query.Limit(meta.Limit).Offset(offset).Find(&images).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Images query failed, this can happened due to database connection failure or database failure", err)
		return
	}

	if len(images) == 0 {
		hp.RespError(ctx, http.StatusNotFound, "Can not found images data", nil, "Data not found in database")
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Images data fetched successfully", images, "", meta)
}
