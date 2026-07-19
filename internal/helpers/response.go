package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponseHeader struct {
	Success    bool        `json:"succes"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	JwtToken   string      `json:"jwt_token,omitempty"`
	Meta       interface{} `json:"meta"`
}

type ErrorResponseHeader struct {
	Success    bool   `json:"succes"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

func RespSuccess(ctx *gin.Context, statusCode int, message string, data interface{}, jwtToken string, meta interface{}) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(statusCode, SuccessResponseHeader{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		JwtToken:   jwtToken,
		Meta:       meta,
	})
}

func RespError(ctx *gin.Context, statusCode int, message string, err error, errStr ...string) {
	finalMessage := message
	if finalMessage == "" {
		finalMessage = http.StatusText(statusCode)
	}

	var finalErrString string

	if len(errStr) > 0 && errStr[0] != "" {
		finalErrString = errStr[0]
	} else if err != nil {
		finalErrString = err.Error()
	}

	ctx.JSON(statusCode, ErrorResponseHeader{
		Success:    false,
		StatusCode: statusCode,
		Message:    finalMessage,
		Error:      finalErrString,
	})
}
