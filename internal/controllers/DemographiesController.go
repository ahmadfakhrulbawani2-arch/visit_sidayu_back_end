package controllers

import (
	"errors"
	"net/http"
	cfg "visit-sidayu-backend/internal/config"
	"visit-sidayu-backend/internal/constants/errorss"
	myE "visit-sidayu-backend/internal/constants/errorss"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/helpers/validation"
	"visit-sidayu-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	singularDem = "Demography"
	pluralDem   = "Demographies"
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
		hp.RespError(ctx, http.StatusInternalServerError, pluralDem+myE.MsgQryErr, err)
		return
	}

	if len(demographies) == 0 {
		hp.RespError(ctx, http.StatusNotFound, singularDem+myE.Msg404Err, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, myE.MsgGet200, demographies, "", meta)
}

func GetDemographyById(ctx *gin.Context) {
	parsedID, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var demography models.Demographies
	err = cfg.DB.First(&demography, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularDem+myE.Msg404Err, err)
			return
		}
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, myE.MsgGet200, demography, "", nil)
}

func CreateDemography(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	input, err := validation.ParseInputJSON[models.CreateDemographies](ctx)
	if err != nil {
		return
	}

	totalPopulation := input.MalePopulation + input.FemalePopulation

	var created models.Demographies
	errCreated := cfg.DB.Where("village_name = ?", input.VillageName).First(&created).Error
	if errCreated == nil {
		hp.RespError(ctx, http.StatusConflict, input.VillageName+myE.Msg409Err, nil, myE.Err409Fill)
		return
	}

	var demography models.Demographies
	copier.Copy(&demography, &input)

	var geo models.Geographies
	err = cfg.DB.Where("village_name = ?", input.VillageName).Select("area", "area_unit").First(&geo).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Village geography data not found", err)
		return
	}

	populationDensity, unit, errPopDen := hp.CalcPopulationDensity(totalPopulation, geo.Area, geo.AreaUnit)
	if errPopDen != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Server failed to calculate population density", err)
		return
	}

	demography.PopulationDensity = populationDensity
	demography.PopulationDensityUnit = unit

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
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	parsedID, err := validation.ParseUrlID(ctx)
	if err != nil {
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
			hp.RespError(ctx, http.StatusNotFound, singularDem+myE.Msg404Err, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hasUpdate := input.VillageName != nil ||
		input.DemographyDataYear != nil ||
		input.MalePopulation != nil ||
		input.FemalePopulation != nil ||
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

	var geo models.Geographies
	err = cfg.DB.Where("village_name = ?", input.VillageName).Select("area", "area_unit").First(&geo).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Village geography data not found", err)
		return
	}

	copier.CopyWithOption(&demography, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	var newTotalPopulation int
	var newPopulationDensity float64
	var unit string
	if input.FemalePopulation != nil || input.MalePopulation != nil {

		malePop := demography.MalePopulation
		if input.MalePopulation != nil {
			malePop = *input.MalePopulation
		}

		femalePop := demography.FemalePopulation
		if input.FemalePopulation != nil {
			femalePop = *input.FemalePopulation
		}

		newTotalPopulation = malePop + femalePop
		newPopulationDensity, unit, err = hp.CalcPopulationDensity(newTotalPopulation, geo.Area, geo.AreaUnit)
		if err != nil {
			hp.RespError(ctx, http.StatusInternalServerError, errorss.MsgSvrErr, err)
		}
	}

	demography.PopulationDensity = newPopulationDensity
	demography.PopulationDensityUnit = unit

	err = cfg.DB.Save(&demography).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularDem+myE.MsgPtc200, demography, "", nil)
}

func DeleteDemography(ctx *gin.Context) {
	// Authorization
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	parsedID, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var demography models.Demographies
	err = cfg.DB.First(&demography, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularDem+myE.Msg404Err, err)
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

	hp.RespSuccess(ctx, http.StatusOK, singularDem+myE.MsgDel200, demography, "", nil)
}

// --- For Geography card & Overview
func GetDistrictDemography(ctx *gin.Context) {

}
