package controllers

import (
	"errors"
	"net/http"
	cfg "visit-sidayu-backend/internal/config"
	hp "visit-sidayu-backend/internal/helpers"
	"visit-sidayu-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GET /api/v1/superadmins?search=john&limit=10&page=1
func GetAllSuperadmins(ctx *gin.Context) {
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	var superadmins []models.Superadmins
	query := cfg.DB.Model(&models.Superadmins{})

	search := ctx.Query("search")
	if search != "" {
		query = query.Where("username ILIKE ? OR email ILIKE ?", search+"%", search+"%")
	}

	meta, offset := hp.CalcMeta(ctx, query)

	err = query.Select("id", "username", "email", "created_at", "updated_at", "deleted_at").Limit(meta.Limit).Offset(offset).Find(&superadmins).Error

	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Superadmins query failed", err)
		return
	}

	if len(superadmins) == 0 {
		hp.RespError(ctx, http.StatusNotFound, "Can not found superadmins data", nil, "Data not found in database")
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Superadmins data fetched successfully", superadmins, "", meta)
}

// GET /api/v1/superadmins/:id
func GetSuperadminByID(ctx *gin.Context) {
	// Authorization
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to parse param 'id'", nil)
		return
	}

	var superadmin models.Superadmins
	err = cfg.DB.First(&superadmin, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Superadmin not found", nil)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch superadmin", err)
		return
	}

	// Jangan kirim password hash ke client
	superadmin.Password = ""

	hp.RespSuccess(ctx, http.StatusOK, "Superadmin fetched successfully", superadmin, "", nil)
}

// PATCH /api/v1/superadmins/:id
func UpdateSuperadmin(ctx *gin.Context) {
	// Authorization
	userID, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Invalid superadmin ID", err)
		return
	}

	if userID != parsedID {
		hp.RespError(
			ctx,
			http.StatusUnauthorized,
			"Unauthorized access",
			nil,
			"Request param 'id' mismatch with current userID",
		)
		return
	}

	var input models.UpdateSuperadminDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming data", err)
		return
	}

	var superadmin models.Superadmins
	err = cfg.DB.First(&superadmin, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Superadmin not found", nil,
				"Database did not find requested superadmin")
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch superadmin", err)
		return
	}

	// Verify current password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(superadmin.Password),
		[]byte(input.Password),
	); err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Current password is incorrect, or you may inserting the same new password with old password", nil)
		return
	}

	updated := false

	// Update username
	if input.Username != nil {
		superadmin.Username = *input.Username
		updated = true
	}

	// Update password
	if input.NewPassword != nil {
		hashedPassword, err := hp.HashPassword(*input.NewPassword)
		if err != nil {
			hp.RespError(ctx, http.StatusInternalServerError, "Failed to hash password", err)
			return
		}

		superadmin.Password = hashedPassword
		updated = true
	}

	if !updated {
		hp.RespError(ctx, http.StatusBadRequest, "No data to update", nil)
		return
	}

	if err := cfg.DB.Save(&superadmin).Error; err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to update superadmin", err)
		return
	}

	// Jangan kirim password hash ke client
	superadmin.Password = ""

	hp.RespSuccess(ctx, http.StatusOK, "Superadmin updated successfully", superadmin, "", nil)
}

// DELETE /api/v1/superadmins/:id
func DeleteSuperadmin(ctx *gin.Context) {
	// Authorization
	userID, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Invalid superadmin ID", err)
		return
	}

	if userID != parsedID {
		hp.RespError(
			ctx,
			http.StatusUnauthorized,
			"Unauthorized access",
			nil,
			"Request param 'id' mismatch with current userID",
		)
		return
	}

	var superadmin models.Superadmins
	err = cfg.DB.First(&superadmin, "id = ?", parsedID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Superadmin not found", nil,
				"Database did not find requested superadmin")
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, "Failed to fetch superadmin", err)
		return
	}

	err = cfg.DB.Unscoped().Delete(&superadmin).Error
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to delete superadmin", err)
		return
	}

	superadmin.Password = ""
	hp.RespSuccess(ctx, http.StatusOK, "Superadmin deleted!", superadmin, "", nil)
}
