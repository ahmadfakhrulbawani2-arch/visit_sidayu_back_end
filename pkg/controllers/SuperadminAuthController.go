package controllers

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"
	cfg "visit-sidayu-backend/pkg/config"
	hp "visit-sidayu-backend/pkg/helpers"
	"visit-sidayu-backend/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	alreadyRegisteredMsg         = "Failed to regist user, user already exist!"
	unknownErrMsg                = "Failed to regist user, unknown internal server error, either BE or db host error, probably because connection error"
	alreadyRegisteredErrEmail    = "Caught re-registration for the same 'email'"
	alreadyRegisteredErrUsername = "Caught re-registration for the same 'username'"
	foundNoSuperadminErrMsg      = "Unexpected not found error, user data not found and probably deleted"
)

// POST /api/v1/superadmins/auth/login
func SuperadminLogin(ctx *gin.Context) {
	var input models.SuperadminsLoginPayload
	if err := ctx.ShouldBindJSON(&input); err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming login input", err)
		return
	}

	var superadmin models.Superadmins
	query := cfg.DB

	// Jika mengandung '@', anggap email.
	// Selain itu anggap username.
	if strings.Contains(input.Identity, "@") {
		query = query.Where("email = ?", input.Identity)
	} else {
		query = query.Where("username = ?", input.Identity)
	}

	if err := query.First(&superadmin).Error; err != nil {
		// mencoba login dan gagal status unauth
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusUnauthorized, foundNoSuperadminErrMsg, err)
			return
		}

		hp.RespError(ctx, http.StatusInternalServerError, unknownErrMsg, err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(superadmin.Password), []byte(input.Password)); err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Failed to login, invalid username/email or password", nil)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": superadmin.ID.String(),
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to create JWT", err)
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Login success!", superadmin, tokenStr, nil)
}

// POST /api/v1/superadmins/auth/register
func SuperAdminRegister(ctx *gin.Context) {
	var input models.CreateSuperadmins
	if err := ctx.ShouldBindJSON(&input); err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming login input", err)
		return
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to encrypt password", err)
		return
	}

	var registeredSuperadmin models.Superadmins
	errEmail := cfg.DB.Where("email = ?", input.Email).First(&registeredSuperadmin).Error
	// email sudah ada
	if errEmail == nil {
		hp.RespError(ctx, http.StatusConflict, alreadyRegisteredMsg, nil, alreadyRegisteredErrEmail)
		return
	}

	errUsername := cfg.DB.Where("username = ?", input.Username).First(&registeredSuperadmin).Error
	// username sudah ada
	if errUsername == nil {
		hp.RespError(ctx, http.StatusConflict, alreadyRegisteredMsg, nil, alreadyRegisteredErrUsername)
		return
	}

	superadmin := models.Superadmins{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPw),
	}

	err = cfg.DB.Create(&superadmin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			hp.RespError(ctx, http.StatusConflict, alreadyRegisteredMsg, err)
		}
		hp.RespError(ctx, http.StatusInternalServerError, unknownErrMsg, err)
		return
	}

	superadmin.Password = ""

	hp.RespSuccess(ctx, http.StatusCreated, "User registered!", superadmin, "", nil)
}

// GET /api/v1/superadmins/auth/me
func SuperadminCurrent(ctx *gin.Context) {
	userID, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	// query
	var user models.Superadmins
	userDataErr := cfg.DB.Select("id", "username", "email", "created_at", "updated_at", "deleted_at").First(&user, "id = ?", userID).Error

	if userDataErr != nil {
		if errors.Is(userDataErr, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusNotFound, "Found no user", nil, foundNoSuperadminErrMsg)
			return
		}
		hp.RespError(ctx, http.StatusInternalServerError, "Database error", nil, userDataErr.Error())
		return
	}

	hp.RespSuccess(ctx, http.StatusOK, "Current user data fetched!", user, "", nil)
}

func GetJwtValidated(ctx *gin.Context) {
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
