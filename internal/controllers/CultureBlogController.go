package controllers

import (
	"errors"
	"net/http"
	cfg "visit-sidayu-backend/internal/config"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func GetAllCultureBlog(ctx *gin.Context) {
	var cBlogs []models.CultureBlogs
	query := cfg.DB.Model(&models.CultureBlogs{})

	search := ctx.Query("search")
	if search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ? OR author ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)
	err := query.Preload("Thumbnail").Preload("Timeline").Preload("Timeline.TimelineData").Limit(meta.Limit).Offset(offset).Find(&cBlogs).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Blogs query failed", err)
		return
	}

	if len(cBlogs) == 0 {
		hp.RespError(ctx, http.StatusNotFound, "Can not found blogs data", nil, "Data not found in database")
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Blogs data fetched successfully", cBlogs, "", meta)
}

func GetCultureBlogById(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to parse param 'id'", nil)
		return
	}

	var cBlog models.CultureBlogs
	err = cfg.DB.Preload("Thumbnail").Preload("Timeline").Preload("Timeline.TimelineData").First(&cBlog, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Blog not found", nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch blog", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Blogs data fetched successfully", cBlog, "", nil)
}

func GetCultureBlogBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	var cBlog models.CultureBlogs
	err := cfg.DB.Preload("Thumbnail").Preload("Timeline").Preload("Timeline.TimelineData").First(&cBlog, "slug = ?", slug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Blog not found", nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch blog", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Blogs data fetched successfully", cBlog, "", nil)
}

func CreateCultureBlog(ctx *gin.Context) {
	// Authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	var input models.CreateCultureBlogsReq
	if err := ctx.ShouldBindJSON(&input); err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming create blog input", err)
		return
	}

	newCBlog := input.ToModel()
	hp.GenerateSlugWithTimestamp(newCBlog.Title, &newCBlog)
	err = cfg.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&newCBlog).Error
	})

	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to insert new blog to database", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusCreated, "New blog created!", newCBlog, "", nil)
}

func UpdateCultureBlog(ctx *gin.Context) {
	// Authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Invalid blog ID", err)
		return
	}

	var input models.UpdateCultureBlogsReq
	if err := ctx.ShouldBindJSON(&input); err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming data", err)
		return
	}

	var cBlog models.CultureBlogs
	err = cfg.DB.First(&cBlog, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Blog not found", nil,
				"Database did not find requested blog")
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch blog", err)
		return
	}

	hasUpdate := input.Title != nil ||
		input.ThemeType != nil ||
		input.Description != nil ||
		input.Content != nil ||
		input.Author != nil ||
		input.BlogWrittenDatetime != nil ||
		input.EstimatedMinutesReadTime != nil ||
		input.ThumbnailID != nil ||
		input.Location != nil ||
		len(input.Tags) > 0 ||
		len(input.ExternalLinks) > 0 ||
		input.Timeline != nil

	if !hasUpdate {
		hp.RespError(ctx, http.StatusBadRequest, "No fields provided for update", nil)
		return
	}

	copier.CopyWithOption(&cBlog, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	if input.Title != nil {
		hp.GenerateSlugWithTimestamp(cBlog.Title, &cBlog)
	}

	if err := cfg.DB.Save(&cBlog).Error; err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to update blog", err)
	}

	hp.RespSuccess(ctx, http.StatusOK, "Blog updated!", cBlog, "", nil)
}

func DeleteCultureBlog(ctx *gin.Context) {
	// Authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Invalid culture blog ID", err)
		return
	}

	var cBlog models.CultureBlogs
	err = cfg.DB.First(&cBlog, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Culture blog not found", nil, "Database did not find requested blog")
			return
		}
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch blog", err)
		return
	}

	err = cfg.DB.Unscoped().Delete(&cBlog).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to delete culture blog", err)
		return
	}
	hp.RespSuccess(ctx, http.StatusOK, "Culture blog deleted!", cBlog, "", nil)
}
