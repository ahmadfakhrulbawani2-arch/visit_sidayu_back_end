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
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	pluralOffc   = "Officials"
	singularOffc = "Official"
)

func GetAllOfficials(ctx *gin.Context) {
	officials := []models.Officials{}
	query := cfg.DB.Model(&models.Officials{})
	search := ctx.Query("search")
	if search != "" {
		query = query.Where("name LIKE ? OR role LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)

	err := query.Preload("Role").Preload("ProfileImage").Offset(offset).Limit(meta.Limit).Find(&officials).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, pluralOffc+myE.MsgQryErr, err)
		return
	}

	if len(officials) == 0 {
		hp.RespError(ctx, http.StatusNotFound, pluralOffc+myE.Msg404Err, nil, myE.Err404Fill)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, pluralOffc+myE.MsgGet200, officials, "", meta)
}

func GetOfficialByID(ctx *gin.Context) {
	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var official models.Officials
	err = cfg.DB.Preload("Role").Preload("ProfileImage").First(&official, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularOffc+myE.Msg404Err, nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularOffc+myE.MsgGet200, official, "", nil)
}

func CreateOfficial(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	input, err := validation.ParseInputJSON[models.CreateOfficials](ctx)
	if err != nil {
		return
	}

	var newOfficial models.Officials
	copier.Copy(&newOfficial, &input)
	err = cfg.DB.Create(&newOfficial).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	if input.ProfileImageID != nil {
		cfg.DB.Model(&newOfficial).Association("ProfileImage").Find(&newOfficial.ProfileImage)
	}

	if input.RoleID != nil {
		cfg.DB.Model(&newOfficial).Association("Role").Find(&newOfficial.Role)
	}

	hp.RespSuccess(ctx, http.StatusCreated, singularOffc+myE.MsgPst201, newOfficial, "", nil)
}

func UpdateOfficial(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	input, err := validation.ParseInputJSON[models.UpdateOfficials](ctx)
	if err != nil {
		return
	}

	var official models.Officials
	err = cfg.DB.First(&official, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularOffc+myE.Msg404Err, err)
			return
		}
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hasUpdate := input.Name != official.Name ||
		input.Description != nil ||
		input.RoleID != nil ||
		input.ProfileImageID != nil

	if !hasUpdate {
		hp.RespError(ctx, http.StatusBadRequest, myE.MsgNoInput, nil, myE.Msg400Err)
		return
	}

	copier.CopyWithOption(&official, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	if input.ProfileImageID != nil {
		if *input.ProfileImageID == uuid.Nil {
			official.ProfileImageID = nil
			official.ProfileImage = nil
		} else {
			val := *input.ProfileImageID
			official.ProfileImageID = &val
		}
	}

	if input.RoleID != nil {
		if *input.RoleID == uuid.Nil {
			official.RoleID = nil
			official.Role = nil
		} else {
			val := *input.RoleID
			official.RoleID = &val
		}
	}

	err = cfg.DB.Session(&gorm.Session{FullSaveAssociations: false}).Save(&official).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.BaseSvrErr, err)
		return
	}

	if official.ProfileImageID != nil {
		cfg.DB.Model(&official).Association("ProfileImage").Find(&official.ProfileImage)
	} else {
		official.ProfileImage = nil
	}

	if official.RoleID != nil {
		cfg.DB.Model(&official).Association("Role").Find(&official.Role)
	} else {
		official.Role = nil
	}

	hp.RespSuccess(ctx, http.StatusOK, singularOffc+myE.MsgPtc200, official, "", nil)
}

func DeleteOfficial(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var official models.Officials
	err = cfg.DB.Preload("Role").Preload("ProfileImage").First(&official, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularOffc+myE.Msg404Err, nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	err = cfg.DB.Unscoped().Delete(&official).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularOffc+myE.MsgDel200, official, "", nil)
}
