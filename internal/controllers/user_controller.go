package controllers

import (
	"errors"
	"net/http"
	"os"
	cfg "simple_go_gin_gorm_postgres_be_template/internal/config"
	"simple_go_gin_gorm_postgres_be_template/internal/helpers"
	res "simple_go_gin_gorm_postgres_be_template/internal/helpers"
	"simple_go_gin_gorm_postgres_be_template/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	alreadyRegisteredMsg = "Failed to regist user, user already exist!"
	unknownErrMsg        = "Failed to regist user, unknown internal server error, either BE or db host error, probably because connection error"
	alreadyRegisteredErr = "Caught re-registration for the same 'email'"
	foundNoUserErrMsg    = "Unexpected not found error, user data not found and probably deleted"
)

type AuthInputRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthInputLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(ctx *gin.Context) {
	var input AuthInputRegister

	// validation
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		res.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming createEvent data", err)
		return
	}

	// hash-password
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if errHash != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Failed to encypt password", err)
		return
	}

	// validation no same email already registered
	var registeredUser models.User
	userDataErr := cfg.DB.Where("email = ?", input.Email).First(&registeredUser).Error

	if userDataErr == nil {
		res.RespError(ctx, http.StatusConflict, alreadyRegisteredMsg, nil, alreadyRegisteredErr)
		return
	}

	if !errors.Is(userDataErr, gorm.ErrRecordNotFound) {
		res.RespError(ctx, http.StatusInternalServerError, unknownErrMsg, userDataErr)
		return
	}

	// passed validation -> save to db
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	// check db writing error
	userCreateErr := cfg.DB.Create(&user).Error
	if userCreateErr != nil {
		if errors.Is(userCreateErr, gorm.ErrDuplicatedKey) {
			res.RespError(ctx, http.StatusConflict, alreadyRegisteredMsg, userCreateErr)
		}
		res.RespError(ctx, http.StatusInternalServerError, unknownErrMsg, userCreateErr)
		return
	}

	res.RespSuccess(ctx, http.StatusCreated, "User registered!", user, "", nil)

}

func UserLogin(ctx *gin.Context) {
	var input AuthInputLogin
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		res.RespError(ctx, http.StatusBadRequest, "Failed to parse incoming createEvent data", err)
		return
	}

	var user models.User
	userDataErr := cfg.DB.Where("email = ?", input.Email).First(&user).Error
	if userDataErr != nil {
		if errors.Is(userDataErr, gorm.ErrRecordNotFound) {
			res.RespError(ctx, http.StatusUnauthorized, "Failed to login, email is not registered", err)
			return
		}
		res.RespError(ctx, http.StatusInternalServerError, "Failed to login user, unknown internal server error, either BE or db host error", userDataErr)
		return
	}

	errMatchPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if errMatchPassword != nil {
		res.RespError(ctx, http.StatusUnauthorized, "Failed to login, input password wrong", errMatchPassword)
		return
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenStr, errToken := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if errToken != nil {
		res.RespError(ctx, http.StatusInternalServerError, "Failed to login, failed to create jwt", userDataErr)
		return
	}

	res.RespSuccess(ctx, http.StatusOK, "Login success!", user, tokenStr, nil)
}

func GetCurrentUser(ctx *gin.Context) {
	userID, err := helpers.GetUserIDFromCtx(ctx)
	if err != nil {
		res.RespError(ctx, http.StatusUnauthorized, "Unauthorized access, read the error value!", err)
		return
	}

	// query
	var user models.User
	userDataErr := cfg.DB.Select("id", "name", "email").First(&user, "id = ?", userID).Error

	if userDataErr != nil {
		if errors.Is(userDataErr, gorm.ErrRecordNotFound) {
			res.RespError(ctx, http.StatusNotFound, "Found no user", nil, foundNoUserErrMsg)
			return
		}
		res.RespError(ctx, http.StatusInternalServerError, "Database error", nil, userDataErr.Error())
		return
	}

	res.RespSuccess(ctx, http.StatusOK, "Current user data fetched!", user, "", nil)
}
