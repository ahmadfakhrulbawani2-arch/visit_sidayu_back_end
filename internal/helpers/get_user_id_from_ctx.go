package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUserIDFromCtx mengambil user_id dari Gin Context dan mem-parsing-nya menjadi uuid.UUID.
func GetUserIDFromCtx(ctx *gin.Context) (uuid.UUID, error) {
	userIdInterface, exists := ctx.Get("user_id")
	if !exists {
		return uuid.Nil, fmt.Errorf("user_id not found in context")
	}

	userIdStr, ok := userIdInterface.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user_id in context is not a string")
	}

	parsedUUID, err := uuid.Parse(userIdStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	return parsedUUID, nil
}
