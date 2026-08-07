package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

// AuditLogUserID logs the user ID for non-GET requests
func AuditLogUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if ctx.Request.Method == "GET" {
			return
		}

		userID, _ := ctx.Get("user_id")

		log.Printf(
			"[AUDIT] user=%v method=%s path=%s status=%d ip=%s",
			userID, ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), ctx.ClientIP(),
		)
	}
}
