package middlewares

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type visitor struct {
	limiter *rate.Limiter
}

var (
	visitors = make(map[string]*visitor)
	mu       sync.Mutex
)

func getVisitor(ip string, r rate.Limit, burst int) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		v = &visitor{
			limiter: rate.NewLimiter(r, burst),
		}
		visitors[ip] = v
	}

	return v.limiter
}

func RateLimiter(limit rate.Limit, burst int) gin.HandlerFunc {
	return func(c *gin.Context) {

		limiter := getVisitor(c.ClientIP(), limit, burst)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Too many requests",
			})
			return
		}

		c.Next()
	}
}
