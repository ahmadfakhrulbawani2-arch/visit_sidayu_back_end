package middlewares

import "github.com/gin-gonic/gin"

func SecurityHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// MIME type sniffing protection
		ctx.Header("X-Content-Type-Options", "nosniff")

		// clickjacking protection
		ctx.Header("X-Frame-Options", "DENY")

		// XSS protection
		ctx.Header("X-XSS-Protection", "1; mode=block")

		// allow other cdn
		ctx.Header("Content-Security-Policy", "")

		// privacy, no reference for other sites
		ctx.Header("Referrer-Policy", "no-referrer")

		ctx.Next()
	}
}
