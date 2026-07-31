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

// GET /api/v1/blogs?search
func GetAllBlogs(ctx *gin.Context) {
	var blogs []models.Blogs
	query := cfg.DB.Model(&models.Blogs{})

	search := ctx.Query("search")
	if search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ? OR author ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)
	err := query.Preload("Thumbnail").Preload("Timeline").Preload("Timeline.TimelineData").Limit(meta.Limit).Offset(offset).Find(&blogs).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Blogs query failed", err)
		return
	}

	if len(blogs) == 0 {
		hp.RespError(ctx, http.StatusInternalServerError, "Blogs"+myE.Msg404Err, nil, myE.Err404Fill)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Blogs data fetched successfully", blogs, "", meta)
}

func GetBlogById(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to parse param 'id'", nil)
		return
	}

	var blog models.Blogs
	err = cfg.DB.Preload("Thumbnail").Preload("Timeline").Preload("Timeline.TimelineData").First(&blog, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Blog not found", nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch blog", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Blogs data fetched successfully", blog, "", nil)
}

func GetBlogBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	var blog models.Blogs
	err := cfg.DB.Preload("Thumbnail").Preload("Timeline").Preload("Timeline.TimelineData").First(&blog, "slug = ?", slug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Blog not found", nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch blog", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Blogs data fetched successfully", blog, "", nil)
}

func CreateBlog(ctx *gin.Context) {
	// Authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	var input models.CreateBlogs
	if err := ctx.ShouldBindJSON(&input); err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming create blog input", err)
		return
	}

	newBlog := input.ToModel()
	hp.GenerateSlugWithTimestamp(newBlog.Title, &newBlog)
	err = cfg.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&newBlog).Error
	})

	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to insert new blog to database", err)
		return
	}

	cfg.DB.Model(&newBlog).Association("Thumbnail").Find(&newBlog.Thumbnail)

	hp.RespSuccess(ctx, http.StatusCreated, "New blog created!", newBlog, "", nil)
}

func UpdateBlog(ctx *gin.Context) {
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

	var input models.UpdateBlogs
	if err := ctx.ShouldBindJSON(&input); err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming data", err)
		return
	}

	var blog models.Blogs
	err = cfg.DB.First(&blog, "id = ?", parsedID).Error
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
		input.Description != nil ||
		input.Content != nil ||
		input.Author != nil ||
		input.BlogWrittenDatetime != nil ||
		input.EstimatedMinutesRead != nil ||
		input.ThumbnailID != nil ||
		input.Location != nil ||
		len(input.Tags) > 0 ||
		len(input.ExternalLinks) > 0 ||
		input.Timeline != nil

	if !hasUpdate {
		hp.RespError(ctx, http.StatusBadRequest, "No fields provided for update", nil)
		return
	}

	copier.CopyWithOption(&blog, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	if input.Title != nil {
		hp.GenerateSlugWithTimestamp(blog.Title, &blog)
	}

	if err := cfg.DB.Save(&blog).Error; err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to update blog", err)
	}

	cfg.DB.Model(&blog).Association("Thumbnail").Find(&blog.Thumbnail)

	hp.RespSuccess(ctx, http.StatusOK, "Blog updated!", blog, "", nil)
}

func DeleteBlog(ctx *gin.Context) {
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

	var blog models.Blogs
	err = cfg.DB.Preload("Thumbnail").Preload("Timeline").Preload("Timeline.TimelineData").First(&blog, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Blog not found", nil,
				"Database did not find requested blog")
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch blog", err)
		return
	}

	err = cfg.DB.Unscoped().Delete(&blog).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to delete superadmin", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Blog deleted!", blog, "", nil)
}
