package middlewares

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newCtx, cancel := context.WithTimeout(ctx.Request.Context(), d)
		defer cancel()

		ctx.Request = ctx.Request.WithContext(newCtx)

		ctx.Next()
	}
}
