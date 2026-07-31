package controllers

import (
	"errors"
	"net/http"
	cfg "visit-sidayu-backend/internal/config"
	myE "visit-sidayu-backend/internal/constants/errorss"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/helpers/validation"
	"visit-sidayu-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	pluralSuBlog   = "Shop and Umkm Blogs"
	singularSuBlog = "Shop and Umkm Blog"
)

func GetAllShopsAndUmkms(ctx *gin.Context) {
	var suBlog []models.ShopsAndUmkmsBlogs
	query := cfg.DB.Model(&models.ShopsAndUmkmsBlogs{})

	search := ctx.Query("search")
	if search != "" {
		query = query.Where("title ILIKE ? OR location ILIKE ? OR marketed_products", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)

	err := query.Preload("Thumbnail").Limit(meta.Limit).Offset(offset).Find(&suBlog).Error

	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, pluralSuBlog+myE.MsgQryErr, err)
		return
	}

	if len(suBlog) == 0 {
		hp.RespError(ctx, http.StatusInternalServerError, pluralSuBlog+myE.Msg404Err, nil, myE.Err404Fill)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, pluralSuBlog+myE.MsgGet200, suBlog, "", meta)
}

func GetShopAndUmkmByID(ctx *gin.Context) {
	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var suBlog models.ShopsAndUmkmsBlogs
	err = cfg.DB.Preload("Thumbnail").First(&suBlog, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularSuBlog+myE.Msg404Err, nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularSuBlog+myE.MsgGet200, suBlog, "", nil)
}

func CreateShopAndUmkm(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	input, err := validation.ParseInputJSON[models.CreateShopsAndUmkmsBlogsReq](ctx)
	if err != nil {
		return
	}

	newSuBlog := input.ToModel()
	hp.GenerateSlugWithTimestamp(newSuBlog.Title, &newSuBlog)
	err = cfg.DB.Create(&newSuBlog).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	if input.ThumbnailID != nil {
		cfg.DB.Model(&newSuBlog).Association("Thumbnail").Find(&newSuBlog.Thumbnail)
	}
}

func UpdateShopAndUmkm(ctx *gin.Context) {

}

func DeleteShopAndUmkm(ctx *gin.Context) {

}
