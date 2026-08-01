package controllers

import (
	cfg "visit-sidayu-backend/internal/config"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// salah desain, wait
func GetAllTimelines(ctx *gin.Context) {
	var timelines []models.Timelines
	query := cfg.DB.Model(&models.Timelines{})
	search := ctx.Query("search")
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)
	err := query.Offset(offset).Limit(meta.Limit).Find(&timelines).Error
	if err != nil {
		return
	}
}

func GetTimelineByID(ctx *gin.Context) {

}

func CreateTimeline(ctx *gin.Context) {

}

func UpdateTimeline(ctx *gin.Context) {

}

func DeleteTimeline(ctx *gin.Context) {

}
