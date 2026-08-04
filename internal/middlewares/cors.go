package middlewares

import (
	"time"
	"visit-sidayu-backend/internal/constants/corscfg"

	"github.com/gin-contrib/cors"
)

// Configures the CORS (Cross-Origin Resource Sharing) middleware
// limiting the allowed sites to those specified in the configuration to access the API
func CorsConfig() cors.Config {
	// only allow this sites to access the API
	allowedOrigins := corscfg.AllowedOrigins
	// only allow these HTTP methods to access the API
	allowedMethods := corscfg.AllowedMethods

	return cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     allowedMethods,
		// only allow these HTTP headers to access the API
		AllowHeaders: []string{
			"Origin",
			"Authorization",
			"Content-Type",
		},
		// only expose these HTTP headers to the client
		ExposeHeaders: []string{
			"Content-Length",
		},
		// allow credentials to be sent with the request
		AllowCredentials: true,
		// cache the preflight CORS configuration for 12 hours
		MaxAge:           12 * time.Hour,
	}
}
