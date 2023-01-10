package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
)

func ProcessTraceID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		format := "X-Trace-Id"
		traceId := ctx.Request.Header.Get(format)

		if traceId == "" {
			u4id, _ := uuid.GenerateUUID()
			traceId = u4id
		}

		ctx.Set(format, traceId)
		ctx.Writer.Header().Set(format, traceId)
		ctx.Next()
	}
}
