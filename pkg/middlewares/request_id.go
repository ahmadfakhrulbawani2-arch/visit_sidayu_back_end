package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// adding request id to the context, so it can be used in the logs
func RequestID() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := uuid.New().String()
		ctx.Set("request_id", id)
		ctx.Writer.Header().Set("X-Request-ID", id)

		ctx.Next()
	}
}
