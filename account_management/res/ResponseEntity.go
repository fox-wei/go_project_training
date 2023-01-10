package res

import (
	"net/http"

	"github.com/fox-wei/go_project_training/life_revive/myerr"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendResponse(ctx *gin.Context, err error, data any) {
	code, message := myerr.DecodeErr(err)

	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
