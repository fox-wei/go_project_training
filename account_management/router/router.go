package router

import (
	"github.com/fox-wei/go_project_training/account_management/handler"
	"github.com/fox-wei/go_project_training/account_management/middleware"
	"github.com/gin-gonic/gin"
)

func Load(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)

	account := engine.Group("account")

	{
		//?添加账户
		account.POST("/:id", handler.AccountHandle.AccountCreate)

		//?删除账户
		account.DELETE("/:id", handler.AccountHandle.AccountDelete1)

		//?更新账户
		account.PUT("/", handler.AccountHandle.AccountUpdate)

		//?根据名字返回账户信息
		account.GET("/:name", handler.AccountHandle.GetAccount)
		//?显示账户
		account.GET("/list", middleware.Auth(), handler.AccountHandle.AccountList)

		//?登录
		account.POST("/login", handler.AccountHandle.Login)

	}

	return engine
}
