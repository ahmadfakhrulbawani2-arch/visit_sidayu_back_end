package validation

import (
	"net/http"
	myE "visit-sidayu-backend/internal/constants/errorss"
	hp "visit-sidayu-backend/internal/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ParseUrlID(ctx *gin.Context) (uuid.UUID, error) {
	id := ctx.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		hp.RespError(ctx, http.StatusInternalServerError, myE.MsgParseParamIdErr, err)
		return uuid.Nil, err
	}

	return parsedID, nil
}
