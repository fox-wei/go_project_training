package main

import (
	"flag"

	"github.com/fox-wei/go_project_training/account_management/handler"
	"github.com/fox-wei/go_project_training/account_management/middleware"
	"github.com/fox-wei/go_project_training/account_management/model"
	"github.com/fox-wei/go_project_training/account_management/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	handler.InItVipe()
	handler.InitDB()
	handler.InitHanlder()
}

func main() {
	flag.Parse()

	model.DB.Init()
	defer model.DB.Close()

	r := gin.New()
	router.Load(r,
		middleware.ProcessTraceID(),
		middleware.Logging())

	r.Run(viper.GetString("addr"))
}
