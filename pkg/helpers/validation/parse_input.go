package validation

import (
	"net/http"
	myE "visit-sidayu-backend/pkg/constants/errorss"
	hp "visit-sidayu-backend/pkg/helpers"

	"github.com/gin-gonic/gin"
)

// ParseInputJSON melakukan binding JSON dan otomatis menangani error response
func ParseInputJSON[T any](ctx *gin.Context) (*T, error) {
	var input T
	if err := ctx.ShouldBindJSON(&input); err != nil {
		hp.RespError(ctx, http.StatusBadRequest, myE.Msg400Err, err)
		return nil, err
	}
	return &input, nil
}
