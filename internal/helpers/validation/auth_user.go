package validation

import (
	"net/http"
	myE "visit-sidayu-backend/internal/constants/errorss"
	hp "visit-sidayu-backend/internal/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthUser(ctx *gin.Context) (uuid.UUID, error) {
	user, err := hp.GetUserIDFromCtx(ctx)
	if err != nil {
		hp.RespError(ctx, http.StatusUnauthorized, myE.Msg401Err, err)
		return uuid.Nil, err
	}

	return user, nil
}
