package controllers

import (
	"math"
	"net/http"
	"simple_go_gin_gorm_postgres_be_template/internal/config"
	cfg "simple_go_gin_gorm_postgres_be_template/internal/config"
	help "simple_go_gin_gorm_postgres_be_template/internal/helpers"
	res "simple_go_gin_gorm_postgres_be_template/internal/helpers"
	"simple_go_gin_gorm_postgres_be_template/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateEvent(ctx *gin.Context) {
	userID, err := help.GetUserIDFromCtx(ctx)
	if err != nil {
		res.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	var payload models.CreateEventRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		res.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming CreateEvent data", err)
		return
	}
	event := models.Event{
		Name:        payload.Name,
		Description: payload.Description,
		Location:    payload.Location,
		DateTime:    payload.DateTime,
		ImageID:     payload.ImageID,
		UserId:      userID,
	}
	if err := config.DB.Create(&event).Error; err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Failed to store data into database, this can happened due to not enough database capacity, database bad connection, etc.", err)
		return
	}

	if err := config.DB.Preload("Image").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "name")
	}).First(&event, "id = ?", event.ID).Error; err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Failed to load image relation metadata", err)
		return
	}

	res.RespSuccess(ctx, http.StatusCreated, "Event created!", event, "", nil)
}

// unprotected
func GetEvents(ctx *gin.Context) {
	var events []models.Event

	query := cfg.DB.Model(&models.Event{})
	search := ctx.Query("search")
	if search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// count total data amount
	var totalRows int64
	query.Count(&totalRows)

	// get query value of page and limit
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

	// calc offset
	offset := (page - 1) * limit

	// calc data / page
	totalPage := int(math.Ceil(float64(totalRows) / float64(limit)))

	err = query.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "email") // don't show password
	}).Preload("Image").Limit(limit).Offset(offset).Find(&events).Error
	if err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Query failed, this can happened due to database connection failure or database failure", err)
		return
	}

	if len(events) == 0 {
		res.RespError(ctx, http.StatusNotFound, "Can not found events data", nil, "Data not found in database")
		return
	}

	meta := models.Meta{
		Page:       page,
		Limit:      limit,
		TotalRows:  totalRows,
		TotalPages: totalPage,
	}
	res.RespSuccess(ctx, http.StatusOK, "Data fetched successfully", events, "", meta)
}

// unprotected
func GetEventById(ctx *gin.Context) {
	var event models.Event
	paramsId := ctx.Param("id")

	parsedParamsId, err := uuid.Parse(paramsId)
	if err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Can not parse parameters", err)
		return
	}

	var errNotFound = config.DB.Preload("Image").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "email")
	}).First(&event, parsedParamsId).Error
	if errNotFound != nil {
		res.RespError(ctx, http.StatusNotFound, "Can not found event data", nil, "Data not found in database")
		return
	}

	res.RespSuccess(ctx, http.StatusOK, "Data fetched successfully", event, "", nil)
}

func UpdateEvent(ctx *gin.Context) {
	userID, err := help.GetUserIDFromCtx(ctx)
	if err != nil {
		res.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	var event models.Event
	paramsId := ctx.Param("id")

	parsedParamsId, err := uuid.Parse(paramsId)
	if err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Can not parse parameters", err)
		return
	}

	var errNotFound = config.DB.First(&event, parsedParamsId).Error
	if errNotFound != nil {
		res.RespError(ctx, http.StatusNotFound, "Can not found event data", nil, "Data not found in database")
		return
	}

	if event.UserId != userID {
		res.RespError(ctx, http.StatusForbidden, "Access denied!", nil, "Caught accessing unauthorized data")
		return
	}

	var input models.Event
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		res.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming UpdateEvent data", err)
		return
	}

	if err := config.DB.Model(&event).Updates(input).Error; err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Failed to store data into database, this can happened due to not enough database capacity, database bad connection, etc.", err)
		return
	}

	if err := config.DB.Preload("Image").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "name")
	}).First(&event, event.ID).Error; err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Failed to store data into database, this can happened due to not enough database capacity, database bad connection, etc.", err)
		return
	}

	res.RespSuccess(ctx, http.StatusOK, "Event data updated!", event, "", nil)
}

func DeleteEvent(ctx *gin.Context) {
	userID, err := help.GetUserIDFromCtx(ctx)
	if err != nil {
		res.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	var event models.Event
	paramsId := ctx.Param("id")

	parsedParamsId, err := uuid.Parse(paramsId)
	if err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Can not parse parameters", err)
	}

	var errNotFound = config.DB.Preload("Image").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "name")
	}).First(&event, parsedParamsId).Error
	if errNotFound != nil {
		res.RespError(ctx, http.StatusNotFound, "Can not found event data", nil, "Data not found in database")
		return
	}

	if event.UserId != userID {
		res.RespError(ctx, http.StatusForbidden, "Access denied!", nil, "Caught accessing unauthorized data")
		return
	}

	if err := config.DB.Unscoped().Delete(&event).Error; err != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Failed to store data into database, this can happened due to not enough database capacity, database bad connection, etc.", err)
		return
	}
	res.RespSuccess(ctx, http.StatusOK, "Event data deleted!", event, "", nil)
}
