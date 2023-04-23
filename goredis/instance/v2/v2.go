package main

import (
	"net/http"
	"path"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//*计算今年已经过去天数和剩余时间
	r := gin.New()
	r.Use(cors.Default())

	_, fileName, _, _ := runtime.Caller(0)
	root := path.Dir(fileName)

	r.Static("/static", root+"/static") //?引入静态文件
	r.LoadHTMLGlob(root + "/view/*")    //&模板文件路径

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "count.html", nil)
	})

	r.GET("/days", func(ctx *gin.Context) {
		//*current time
		t := time.Now()

		//*computer passed days
		yearStart := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.Local)
		daysPassed := int(t.Sub(yearStart).Hours() / 24)

		//*computer the left days
		newYear := time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, time.Local)
		daysLeft := int(newYear.Sub(t).Hours() / 24)

		//*current time
		current := t.Format("2006年01月02日")

		ctx.JSON(http.StatusOK, gin.H{
			"daysPassed": daysPassed,
			"daysLeft":   daysLeft,
			"current":    current,
		})
	})

	r.Run()
}
