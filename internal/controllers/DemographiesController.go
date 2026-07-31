package controllers

import (
	"errors"
	"net/http"
	cfg "visit-sidayu-backend/internal/config"
	"visit-sidayu-backend/internal/constants/errorss"
	myE "visit-sidayu-backend/internal/constants/errorss"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/helpers/mtk"
	"visit-sidayu-backend/internal/helpers/validation"
	"visit-sidayu-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	singularDem = "Demography"
	pluralDem   = "Demographies"
	distDem     = "District demography"
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
	err = cfg.DB.Where("village_name = ?", input.VillageName).First(&geo).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
	demography.TotalPopulation = totalPopulation

	// demography should be using Save as it was created when user create geographies first
	err = cfg.DB.Save(&demography).Error
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
		hp.RespError(ctx, http.StatusBadRequest, myE.MsgNoInput, nil, myE.Msg400Err)
		return
	}

	var geo models.Geographies
	err = cfg.DB.Where("village_name = ?", demography.VillageName).First(&geo).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Village geography data not found", err)
		return
	}

	// make sure to copy after fetch the geography
	copier.CopyWithOption(&demography, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	var newTotalPopulation int
	var newPopulationDensity float64
	var unit *string
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
	demography.TotalPopulation = newTotalPopulation

	// in case there's name change
	geo.VillageName = demography.VillageName

	err = cfg.DB.Save(&geo).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

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
	var data models.GetDistrictDemographies

	// dbD := cfg.DB.Model(&models.Demographies{})
	// dbG := cfg.DB.Model(&models.Geographies{})

	err := cfg.DB.Model(&models.Demographies{}).Select("demography_data_year").Group("demography_data_year").Order("count(*) DESC").Limit(1).Scan(&data.DemographyDataYear).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, distDem+myE.MsgQryErr, err)
		return
	}

	err = cfg.DB.Model(&models.Demographies{}).Select("SUM(male_population)").Scan(&data.MalePopulation).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, distDem+myE.MsgQryErr, err)
		return
	}

	err = cfg.DB.Model(&models.Demographies{}).Select("SUM(female_population)").Scan(&data.FemalePopulation).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, distDem+myE.MsgQryErr, err)
		return
	}

	err = cfg.DB.Model(&models.Demographies{}).Select("SUM(total_population)").Scan(&data.TotalPopulation).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, distDem+myE.MsgQryErr, err)
		return
	}

	// density pop from tot pop / tot area
	var (
		distAreaKm2 float64
		distAreaHa  float64
	)
	err = cfg.DB.Model(&models.Geographies{}).Select("COALESCE(SUM(area), 0)").Where("area_unit = ?", "km2").Scan(&distAreaKm2).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, distDem+myE.MsgQryErr, err)
		return
	}

	err = cfg.DB.Model(&models.Geographies{}).Select("COALESCE(SUM(area), 0)").Where("area_unit = ?", "ha").Scan(&distAreaHa).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, distDem+myE.MsgQryErr, err)
		return
	}

	distAreaKm2 += mtk.SwitchHaToKm2(distAreaHa)

	data.PopulationDensity = float64(data.TotalPopulation) / distAreaKm2
	data.PopulationDensityUnit = "jiwa/km2"

	var sources []string
	err = cfg.DB.Model(&models.Demographies{}).Distinct("source_name").Pluck("source_name", &sources).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, distDem+myE.MsgQryErr, err)
		return
	}

	data.Sources = pq.StringArray(sources)

	hp.RespSuccess(ctx, http.StatusOK, distDem+myE.MsgGet200, data, "", nil)
}
