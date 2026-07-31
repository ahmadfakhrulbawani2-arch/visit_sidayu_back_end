package controllers

import (
	"errors"
	"net/http"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/models"

	cfg "visit-sidayu-backend/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	// "visit-sidayu-backend/internal/constants/errorss"
	myE "visit-sidayu-backend/internal/constants/errorss"
	// hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/helpers/validation"
)

var (
	singularGal = "Gallery"
	pluralGal   = "Galleries"
)

func GetAllGalleries(ctx *gin.Context) {
	var galleries []models.Galleries
	query := cfg.DB.Model(&models.Galleries{})
	search := ctx.Query("search")
	if search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)
	err := query.Preload("Image").Limit(meta.Limit).Offset(offset).Find(&galleries).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, pluralGal+myE.MsgQryErr, err)
		return
	}

	if len(galleries) == 0 {
		hp.RespError(ctx, http.StatusNotFound, pluralGal+myE.Msg404Err, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, pluralGal+myE.MsgGet200, galleries, "", meta)
}

func GetGalleryByID(ctx *gin.Context) {
	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var gallery models.Galleries
	err = cfg.DB.Preload("Image").First(&gallery, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularGal+myE.Msg404Err, err)
			return
		}
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, myE.MsgGet200, gallery, "", nil)
}

func GetGalleryBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	var gallery models.Galleries
	err := cfg.DB.Preload("Image").First(&gallery, "slug = ?", slug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularGal+myE.Msg404Err, err)
			return
		}
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, myE.MsgGet200, gallery, "", nil)
}

func CreateGallery(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	input, err := validation.ParseInputJSON[models.CreateGalleries](ctx)
	if err != nil {
		return
	}

	var newGallery models.Galleries
	copier.Copy(&newGallery, &input)
	hp.GenerateSlugWithTimestamp(newGallery.Name, &newGallery)
	err = cfg.DB.Create(&newGallery).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.BaseSvrErr, err)
		return
	}

	cfg.DB.Model(&newGallery).Association("Image").Find(&newGallery.Image)

	hp.RespSuccess(ctx, http.StatusCreated, myE.MsgPst201, newGallery, "", nil)
}

func UpdateGallery(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	input, err := validation.ParseInputJSON[models.UpdateGalleries](ctx)
	if err != nil {
		return
	}

	var gallery models.Galleries
	err = cfg.DB.First(&gallery, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Blog not found", nil,
				"Database did not find requested blog")
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch blog", err)
		return
	}

	hasUpdate := input.Name != nil ||
		input.Description != nil ||
		input.ImageID != nil

	if !hasUpdate {
		hp.RespError(ctx, http.StatusBadRequest, myE.MsgNoInput, nil, myE.Msg400Err)
		return
	}

	copier.CopyWithOption(&gallery, input, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	if input.Name != nil {
		hp.GenerateSlugWithTimestamp(gallery.Name, &gallery)
	}

	err = cfg.DB.Save(&gallery).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	cfg.DB.Model(&gallery).Association("Image").Find(&gallery.Image)

	hp.RespSuccess(ctx, http.StatusOK, singularGal+myE.MsgPtc200, gallery, "", nil)

}

func DeleteGallery(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var gallery models.Galleries
	err = cfg.DB.Preload("Image").First(&gallery, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularDem+myE.Msg404Err, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	err = cfg.DB.Unscoped().Delete(&gallery).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularGal+myE.MsgDel200, gallery, "", nil)
}
