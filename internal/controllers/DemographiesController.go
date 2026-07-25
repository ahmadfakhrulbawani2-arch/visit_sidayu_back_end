package controllers

import (
	"errors"
	"net/http"
	cfg "visit-sidayu-backend/internal/config"
	myE "visit-sidayu-backend/internal/constants/errorss"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	singularName = "Demography"
	pluralName   = "Demographies"
)

// --- Per row
func GetAllDemographies(ctx *gin.Context) {
	var demographies []models.Demographies
	query := cfg.DB.Model(&models.Demographies{})
	search := ctx.Query("search")
	if search != "" {
		query = query.Where("village_name ILIKE ?", "%"+search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)
	err := query.Limit(meta.Limit).Offset(offset).Find(&demographies).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, pluralName+myE.MsgQryErr, err)
		return
	}

	if len(demographies) == 0 {
		hp.RespError(ctx, http.StatusInternalServerError, singularName+myE.Msg404Err, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, myE.MsgGet200, demographies, "", meta)
}

func GetDemographyById(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgParseParamIdErr, err)
		return
	}

	var demography models.Demographies
	err = cfg.DB.First(&demography, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularName+myE.Msg404Err, err)
			return
		}
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, myE.MsgGet200, demography, "", nil)
}

func CreateDemography(ctx *gin.Context) {
	// Authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, myE.Msg401Err, err)
		return
	}

	var input models.CreateDemographies
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		hp.RespError(ctx, http.StatusBadRequest, myE.Msg400Err, err)
		return
	}

	var created models.Demographies
	errCreated := cfg.DB.Where("village_name = ?", input.VillageName).First(&created).Error
	if errCreated == nil {
		hp.RespError(ctx, http.StatusConflict, input.VillageName+myE.Msg409Err, nil, myE.Err409Fill)
		return
	}

	var demography models.Demographies
	copier.Copy(&demography, &input)

	err = cfg.DB.Create(&demography).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			hp.RespError(ctx, http.StatusConflict, "Some key are"+myE.Msg409Err, err)
			return
		}
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, myE.MsgPst201, demography, "", nil)
}

func UpdateDemography(ctx *gin.Context) {
	// Authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, myE.Msg401Err, err)
		return
	}

	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Invalid blog ID", err)
		return
	}

	var input models.UpdateDemographies
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		hp.RespError(ctx, http.StatusBadRequest, myE.Msg400Err, err)
		return
	}

	var demography models.Demographies
	err = cfg.DB.First(&demography, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularName+myE.Msg404Err, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hasUpdate := input.VillageName != nil ||
		input.DemographyDataYear != nil ||
		input.MalePopulation != nil ||
		input.FemalePopulation != nil ||
		input.TotalPopulation != nil ||
		input.PopulationDensityUnit != nil ||
		input.FamiliesNumber != nil ||
		input.NumberOfBirth != nil ||
		input.NumberOfDeath != nil ||
		input.WorkingPopulation != nil ||
		input.UnemployedPopulation != nil ||
		input.HousekeepingPopulation != nil ||
		input.StudentPopulation != nil ||
		input.SourceName != nil ||
		input.ExternalLinkSource != nil

	if !hasUpdate {
		hp.RespError(ctx, http.StatusBadRequest, myE.MsgNoInput, nil)
		return
	}

	copier.CopyWithOption(&demography, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	err = cfg.DB.Save(&demography).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularName+myE.MsgPtc200, demography, "", nil)
}

func DeleteDemography(ctx *gin.Context) {
	// Authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, myE.Msg401Err, err)
		return
	}

	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Invalid blog ID", err)
		return
	}

	var demography models.Demographies
	err = cfg.DB.First(&demography, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularName+myE.Msg404Err, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	err = cfg.DB.Unscoped().Delete(&demography).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularName+myE.MsgDel200, demography, "", nil)
}

// --- For Geography card & Overview
