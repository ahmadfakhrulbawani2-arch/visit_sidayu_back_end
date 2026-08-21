package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(superadminID string) (accessToken string, refreshToken string, err error) {
	// 1. Access Token (Berumur pendek, misal 1 jam)
	accessClaims := jwt.MapClaims{
		"sub":  superadminID,
		"type": "access",
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	}
	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = accessTokenObj.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	// 2. Refresh Token (Berumur panjang, misal 7 hari)
	refreshClaims := jwt.MapClaims{
		"sub":  superadminID,
		"type": "refresh",
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshTokenObj.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET"))) // Bisa pakai secret khusus refresh atau JWT_SECRET yang sama
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
