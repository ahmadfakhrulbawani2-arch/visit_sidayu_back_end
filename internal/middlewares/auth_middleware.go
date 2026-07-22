package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequiredAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":     false,
				"status_code": http.StatusUnauthorized,
				"message":     "Access denied, token is not found or invalid format",
				"error":       "Unauthorized action detected, Authorization token with 'Bearer' prefix is not found",
			})
			return
		}
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		token, errTokenStr := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if errTokenStr != nil {
			fmt.Println(errTokenStr)
			fmt.Println(os.Getenv("JWT_SECRET"))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("sa_id", claims["sub"].(string))
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":     false,
				"status_code": http.StatusUnauthorized,
				"message":     "Access denied, token is valid by format but failed in JWT validation, the token is probably expired, try to relogin",
				"error":       "Unauthorized action detected, JWT validation failed",
			})
			return
		}
	}
}
