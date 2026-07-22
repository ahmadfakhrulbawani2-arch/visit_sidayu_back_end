package helpers

import (
	"math"
	"strconv"
	"visit-sidayu-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CalcMeta(ctx *gin.Context, db *gorm.DB) (meta models.Meta, rowOffset int) {
	var totalRows int64
	db.Session(&gorm.Session{}).Count(&totalRows)

	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var totalPages int
	if limit > 0 {
		totalPages = int(math.Ceil(float64(totalRows) / float64(limit)))
	}

	meta = models.Meta{
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}

	return meta, offset
}
