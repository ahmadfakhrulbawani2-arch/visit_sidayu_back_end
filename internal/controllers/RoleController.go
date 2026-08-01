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
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	pluralRole   = "Roles"
	singularRole = "Role"
)

func GetAllRoles(ctx *gin.Context) {
	var roles []models.Roles
	query := cfg.DB.Model(&models.Roles{})
	search := ctx.Query("search")
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)

	err := query.Limit(meta.Limit).Offset(offset).Find(&roles).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, pluralRole+myE.MsgQryErr, err)
		return
	}

	if len(roles) == 0 {
		hp.RespError(ctx, http.StatusNotFound, pluralRole+myE.Msg404Err, nil, myE.Err404Fill)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, pluralRole+myE.MsgGet200, roles, "", meta)
}

func GetRoleById(ctx *gin.Context) {
	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var role models.Roles
	err = cfg.DB.First(&role, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularRole+myE.Msg404Err, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularRole+myE.MsgGet200, role, "", nil)
}

func CreateRole(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	input, err := validation.ParseInputJSON[models.CreateRoles](ctx)
	if err != nil {
		return
	}

	var newRole models.Roles
	copier.Copy(&newRole, &input)
	err = cfg.DB.Create(&newRole).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusCreated, singularRole+myE.MsgPst201, newRole, "", nil)
}

func UpdateRole(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	input, err := validation.ParseInputJSON[models.UpdateRoles](ctx)
	if err != nil {
		return
	}

	if input.Name != nil && *input.Name == "" {
		hp.RespError(ctx, http.StatusBadRequest, myE.Msg400Err, nil, "name tidak boleh kosong")
		return
	}

	if input.Level != nil && *input.Level == 0 {
		hp.RespError(ctx, http.StatusBadRequest, myE.Msg400Err, nil, "level tidak boleh 0")
		return
	}

	if input.Description != nil && *input.Description == "" {
		hp.RespError(ctx, http.StatusBadRequest, myE.Msg400Err, nil, "description tidak boleh kosong")
		return
	}

	var role models.Roles
	err = cfg.DB.First(&role, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularRole+myE.Msg404Err, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	// this is actually redundant because we already checked for input validation above, but yeah it's good to have
	hasUpdate := *input.Name != role.Name || *input.Level != role.Level || *input.Description != role.Description
	if !hasUpdate {
		hp.RespError(ctx, http.StatusBadRequest, myE.MsgNoInput, nil, myE.Err404Fill)
		return
	}

	copier.CopyWithOption(&role, &input, copier.Option{IgnoreEmpty: true})

	err = cfg.DB.Save(&role).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularRole+myE.MsgPtc200, role, "", nil)
}

func DeleteRole(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var role models.Roles
	err = cfg.DB.First(&role, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularRole+myE.Msg404Err, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	err = cfg.DB.Unscoped().Delete(&role).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularRole+myE.MsgPtc200, role, "", nil)
}
