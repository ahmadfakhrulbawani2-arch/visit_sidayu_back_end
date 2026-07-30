package controllers

import (
	"errors"
	"net/http"
	"slices"
	cfg "visit-sidayu-backend/internal/config"
	"visit-sidayu-backend/internal/constants/errorss"
	myE "visit-sidayu-backend/internal/constants/errorss"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/helpers/validation"
	"visit-sidayu-backend/internal/models"

	inCfg "visit-sidayu-backend/internal/constants/input"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	singularGeo = "Geography"
	pluralGeo   = "Geographies"
)

func GetAllGeographies(ctx *gin.Context) {
	var geographies []models.Geographies
	query := cfg.DB.Model(&models.Geographies{})
	search := ctx.Query("search")
	if search != "" {
		query = query.Where("village_name ILIKE ?", "%"+search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)

	err := query.Limit(meta.Limit).Offset(offset).Find(&geographies).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, pluralGeo+myE.MsgQryErr, err)
		return
	}

	if len(geographies) == 0 {
		hp.RespError(ctx, http.StatusNotFound, singularGeo+myE.Msg404Err, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, myE.MsgGet200, geographies, "", meta)
}

func GetGeographyByID(ctx *gin.Context) {
	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var geography models.Geographies
	err = cfg.DB.First(&geography, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularGeo+myE.Msg404Err, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, myE.MsgGet200, geography, "", nil)
}

func CreateGeography(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	// input validation
	input, err := validation.ParseInputJSON[models.CreateGeographies](ctx)
	if err != nil {
		return
	}

	// check area unit
	if !slices.Contains(inCfg.Area_Units, input.AreaUnit) {
		hp.RespError(ctx, http.StatusBadRequest, myE.Msg400Err, err)
		return
	}

	var created models.Geographies
	err = cfg.DB.Where("village_name = ?", input.VillageName).First(&created).Error
	if err == nil {
		hp.RespError(ctx, http.StatusConflict, input.VillageName+myE.Msg409Err, nil, myE.Err409Fill)
		return
	}

	var geography models.Geographies
	copier.Copy(&geography, &input)

	// update demography if found
	var demography models.Demographies
	err = cfg.DB.Where("village_name = ?", geography.VillageName).First(&demography).Error

	notFound := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !notFound {
		hp.RespError(ctx, http.StatusInternalServerError, "Village geography data not found", err)
		return
	}

	demography.PopulationDensity, demography.PopulationDensityUnit, err = hp.CalcPopulationDensity(demography.TotalPopulation, geography.Area, geography.AreaUnit)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, errorss.MsgSvrErr, err)
	}

	if !notFound {
		// need to make it only update if found, to prevent creating unwanted record
		err = cfg.DB.Save(&demography).Error
		if err != nil {
			hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
			return
		}
	}

	// then create the geography

	err = cfg.DB.Create(&geography).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			hp.RespError(ctx, http.StatusConflict, "Some key are"+myE.Msg409Err, err)
			return
		}
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, myE.MsgPst201, geography, "", nil)
}

func UpdateGeography(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	// id url
	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	// input validation
	input, err := validation.ParseInputJSON[models.UpdateGeographies](ctx)
	if err != nil {
		return
	}

	if !slices.Contains(inCfg.Area_Units, *input.AreaUnit) {
		hp.RespError(ctx, http.StatusBadRequest, myE.Msg400Err, err)
		return
	}

	var geography models.Geographies
	err = cfg.DB.First(&geography, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularGeo+myE.Msg404Err, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hasUpdate := input.Area != nil ||
		input.AreaUnit != nil ||
		input.ImageID != nil ||
		input.RainfallRate != nil ||
		input.RainyDay != nil ||
		input.Source != nil ||
		input.VillageName != nil

	if !hasUpdate {
		hp.RespError(ctx, http.StatusBadRequest, myE.MsgNoInput, nil)
		return
	}

	copier.CopyWithOption(&geography, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// fetch related demography data to update it
	var demography models.Demographies
	err = cfg.DB.Where("village_name = ?", geography.VillageName).First(&demography).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Village demography data not found", err)
		return
	}

	// calc new pop den
	newPopulationDensity, unit, err := hp.CalcPopulationDensity(demography.TotalPopulation, geography.Area, geography.AreaUnit)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, errorss.MsgSvrErr, err)
	}

	demography.PopulationDensity = newPopulationDensity
	demography.PopulationDensityUnit = unit

	// save it all the way
	err = cfg.DB.Save(&demography).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	err = cfg.DB.Save(&geography).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularGeo+myE.MsgPtc200, geography, "", nil)
}

func DeleteGeograhy(ctx *gin.Context) {
	// Authorization
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var geography models.Geographies
	err = cfg.DB.First(&geography, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularGeo+myE.Msg404Err, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	err = cfg.DB.Unscoped().Delete(&geography).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularGeo+myE.MsgDel200, geography, "", nil)
}
