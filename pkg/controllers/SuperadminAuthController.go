package controllers

import (
	"errors"
	"net/http"
	"os"
	"strings"
	cfg "visit-sidayu-backend/pkg/config"
	"visit-sidayu-backend/pkg/helpers"
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

	accessToken, refreshToken, err := helpers.GenerateTokens(superadmin.ID.String())
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to create tokens", err)
		return
	}

	isSecure := os.Getenv("GIN_MODE") == "production" // True jika production
	ctx.SetCookie(
		"refresh_token",    // Nama cookie
		refreshToken,       // Nilai refresh token
		7*24*60*60,         // Max age dalam detik (7 hari)
		"/api/v1/superadmins/auth/jwt/refresh", // refresh token endpoint
		"",                 // Domain
		isSecure,           // Secure (true jika HTTPS)
		true,               // HttpOnly (mencegah akses via JavaScript / XSS)
	)

	hp.RespSuccess(ctx, http.StatusOK, "Login success!", superadmin, accessToken, nil)
}

// POST /api/v1/superadmins/auth/register
func SuperAdminRegister(ctx *gin.Context) {
	_, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}
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

// POST /api/v1/superadmins/auth/jwt
func PostGetJwtValidated(ctx *gin.Context) {
	userID, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	var input models.GetSuperadminJwtValidated
	if err := ctx.ShouldBindJSON(&input); err != nil {
		hp.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming login input", err)
		return
	}

	var user models.Superadmins
	userDataErr := cfg.DB.Select("username", "email").First(&user, "id = ?", userID).Error

	if userDataErr != nil {
		if errors.Is(userDataErr, gorm.ErrRecordNotFound) {
			hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
			return
		}
		hp.RespError(ctx, http.StatusInternalServerError, "Database error", nil, userDataErr.Error())
		return
	}

	if user.Username != input.Username || user.Email != input.Email {
		hp.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", nil, "username or email does not match")
		return
	}

	ctx.Status(http.StatusNoContent)
}

// POST /api/v1/superadmins/auth/jwt/refresh
func SuperadminRefreshToken(ctx *gin.Context) {
	// 🍪 Ambil Refresh Token dari Cookie
	cookieRefreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "Refresh token not found in cookie", err)
		return
	}

	// Parse dan validasi Refresh Token
	token, err := jwt.Parse(cookieRefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil // Atau os.Getenv("JWT_SECRET")
	})

	if err != nil || !token.Valid {
		hp.RespError(ctx, http.StatusUnauthorized, "Invalid or expired refresh token", err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		hp.RespError(ctx, http.StatusUnauthorized, "Invalid token claims", nil)
		return
	}

	// Pastikan tipe token adalah "refresh"
	if tokenType, ok := claims["type"].(string); !ok || tokenType != "refresh" {
		hp.RespError(ctx, http.StatusUnauthorized, "Invalid token type", nil)
		return
	}

	superadminID, ok := claims["sub"].(string)
	if !ok {
		hp.RespError(ctx, http.StatusUnauthorized, "Invalid token subject", nil)
		return
	}

	// Cek apakah user masih ada di database
	var superadmin models.Superadmins
	if err := cfg.DB.First(&superadmin, "id = ?", superadminID).Error; err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, "User no longer exists", err)
		return
	}

	// Generate sepasang token baru (opsional: implementasikan token rotation)
	newAccessToken, newRefreshToken, err := helpers.GenerateTokens(superadminID)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, "Failed to generate new tokens", err)
		return
	}

	// 🔄 Perbarui Refresh Token di Cookie (Token Rotation)
	isSecure := os.Getenv("GIN_MODE") == "release"
	ctx.SetCookie(
		"refresh_token",
		newRefreshToken,
		7*24*60*60,
		"/api/v1/superadmins/auth/refresh",
		"",
		isSecure,
		true,
	)

	// Kembalikan Access Token yang baru melalui JSON response body
	responseData := map[string]any{
		"access_token": newAccessToken,
	}

	hp.RespSuccess(ctx, http.StatusOK, "Token refreshed successfully!", responseData, "", nil)
}

// POST /api/v1/superadmins/auth/jwt/logout
func SuperadminLogout(ctx *gin.Context) {
	isSecure := os.Getenv("GIN_MODE") == "release"
	ctx.SetCookie(
		"refresh_token",
		"",
		-1,
		"/api/v1/superadmins/auth/jwt/refresh", // Sesuaikan path-nya agar sama persis dengan saat set cookie
		"",
		isSecure,
		true,
	)

	// hp.RespSuccess(ctx, http.StatusOK, "Logout success!", nil, "", nil)
	ctx.Status(http.StatusNoContent)
}
