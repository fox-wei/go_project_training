package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	mylog "github.com/fox-wei/go_project_training/account_management/MyLog"
	"github.com/fox-wei/go_project_training/account_management/model/log"
	"github.com/fox-wei/go_project_training/account_management/myerr"
	"github.com/fox-wei/go_project_training/account_management/res"
	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
)

func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now().UTC()
		path := ctx.Request.URL.Path

		reqUrl := ctx.Request.URL.Path
		if strings.Contains(reqUrl, "image") {
			return
		}

		var bodies []byte
		if ctx.Request.Body != nil {
			bodies, _ = ioutil.ReadAll(ctx.Request.Body)
		}

		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodies))
		method := ctx.Request.Method
		ip := ctx.ClientIP()

		blw := &log.BodyLoggerWriter{
			ResponseWriter: ctx.Writer,
			Body:           bytes.NewBufferString(""),
		}

		ctx.Writer = blw

		ctx.Next() //!执行handler

		end := time.Now().UTC()
		sub := end.Sub(start)
		code := -100
		message := ""
		var response res.Response

		if err := json.Unmarshal(blw.Body.Bytes(), &response); err != nil {
			code = myerr.InternalServerError.Code
			message = err.Error()
			mylog.Log.Error(err)
		} else {
			code = response.Code
			message = response.Message
		}

		mylog.Log.Infof("%-13s | %-12s | %s %s {code: %d, message: %s}", sub, ip, pad.Right(method, 5, ""), path, code, message)
	}

}
