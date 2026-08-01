package controllers

import (
	"errors"
	"net/http"
	"slices"
	cfg "visit-sidayu-backend/internal/config"
	myE "visit-sidayu-backend/internal/constants/errorss"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/helpers/validation"
	"visit-sidayu-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	pluralIBlogs  = "Industries blogs"
	singularIBlog = "Industry blog"
)

func GetAllIndustriesBlog(ctx *gin.Context) {
	var iBlog []models.IndustriesBlogs
	query := cfg.DB.Model(&models.IndustriesBlogs{})

	search := ctx.Query("search")
	if search != "" {
		query = query.Where("title ILIKE ? OR location ILIKE ? OR produced_products ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)

	err := query.Preload("Thumbnail").Limit(meta.Limit).Offset(offset).Find(&iBlog).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, pluralIBlogs+myE.MsgQryErr, err)
		return
	}

	if len(iBlog) == 0 {
		hp.RespError(ctx, http.StatusInternalServerError, pluralIBlogs+myE.Msg404Err, nil, myE.Err404Fill)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, pluralIBlogs+myE.MsgGet200, iBlog, "", meta)
}

func GetIndustryBlogByID(ctx *gin.Context) {
	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var iBlog models.IndustriesBlogs
	err = cfg.DB.Preload("Thumbnail").First(&iBlog, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularIBlog+myE.Msg404Err, nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularIBlog+myE.MsgGet200, iBlog, "", nil)
}

func GetIndustryBlogBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	var iBlog models.IndustriesBlogs
	err := cfg.DB.Preload("Thumbnail").First(&iBlog, "slug = ?", slug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularIBlog+myE.Msg404Err, nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularIBlog+myE.MsgGet200, iBlog, "", nil)
}

func CreateIndustryBlog(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	input, err := validation.ParseInputJSON[models.CreateIndustriesBlogsReq](ctx)
	if err != nil {
		return
	}

	newIBlog := input.ToModel()
	hp.GenerateSlugWithTimestamp(newIBlog.Title, &newIBlog)
	err = cfg.DB.Create(&newIBlog).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	if input.ThumbnailID != nil {
		cfg.DB.Model(&newIBlog).Association("Thumbnail").Find(&newIBlog.Thumbnail)
	}

	hp.RespSuccess(ctx, http.StatusCreated, singularIBlog+myE.MsgPst201, newIBlog, "", nil)

}

func UpdateIndustryBlog(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	input, err := validation.ParseInputJSON[models.UpdateIndustriesBlogsReq](ctx)
	if err != nil {
		return
	}

	var iBlog models.IndustriesBlogs
	err = cfg.DB.Preload("Thumbnail").First(&iBlog, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularIBlog+myE.Msg404Err, nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hasUpdate := input.Title != iBlog.Title ||
		input.Content != iBlog.Content ||
		input.Location != nil ||
		input.Rating != nil ||
		input.Revenue != nil ||
		(input.ProducedProducts != nil && !slices.Equal(input.ProducedProducts, iBlog.ProducedProducts)) ||
		input.ProductionRatesPiecePerDay != nil ||
		input.ThumbnailID != iBlog.ThumbnailID ||
		input.YearFounded != iBlog.YearFounded ||
		input.EmployeesCount != iBlog.EmployeesCount ||
		input.BusinessType != iBlog.BusinessType

	if !hasUpdate {
		hp.RespError(ctx, http.StatusBadRequest, myE.MsgNoInput, nil, myE.Msg400Err)
		return
	}

	if input.Title != iBlog.Title {
		hp.GenerateSlugWithTimestamp(input.Title, &iBlog)
	}

	copier.CopyWithOption(&iBlog, &input, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// 2. Tangani ThumbnailID secara manual agar bisa jadi NULL atau terganti dengan benar
	if input.ThumbnailID != nil {
		if *input.ThumbnailID == uuid.Nil {
			// Jika dikirim null uuid (00000000-0000-0000-0000-000000000000), set NULL di DB
			iBlog.ThumbnailID = nil
			iBlog.Thumbnail = nil // Kosongkan relasi struct-nya juga
		} else {
			// Jika dikirim UUID valid, assign nilainya
			val := *input.ThumbnailID
			iBlog.ThumbnailID = &val
		}
	}

	// Gunakan Save / Updates GORM.
	// Catatan: Jika menggunakan .Save(), pastikan field iBlog.ThumbnailID terupdate dengan benar ke database.
	err = cfg.DB.Session(&gorm.Session{FullSaveAssociations: false}).Save(&iBlog).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.BaseSvrErr, err)
		return
	}

	// Reload thumbnail jika ada
	if iBlog.ThumbnailID != nil {
		cfg.DB.Model(&iBlog).Association("Thumbnail").Find(&iBlog.Thumbnail)
	} else {
		iBlog.Thumbnail = nil
	}

	hp.RespSuccess(ctx, http.StatusOK, singularIBlog+myE.MsgPtc200, iBlog, "", nil)

}

func DeleteIndustryBlog(ctx *gin.Context) {
	_, err := validation.AuthUser(ctx)
	if err != nil {
		return
	}

	id, err := validation.ParseUrlID(ctx)
	if err != nil {
		return
	}

	var iBlog models.IndustriesBlogs
	err = cfg.DB.Preload("Thumbnail").First(&iBlog, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, singularIBlog+myE.Msg404Err, nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	err = cfg.DB.Unscoped().Delete(&iBlog).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgSvrErr, err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, singularIBlog+myE.MsgDel200, iBlog, "", nil)
}
