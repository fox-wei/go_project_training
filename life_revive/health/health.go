package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//?result "OK"
func Health(ctx *gin.Context) {
	message := "OK"
	ctx.JSON(http.StatusOK, "\n"+message)
}
