package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// limit the body size to prevent memory exhaustion
func BodyLimit(limit int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.Body = http.MaxBytesReader(
			ctx.Writer,
			ctx.Request.Body,
			limit,
		)

		ctx.Next()
	}
}
