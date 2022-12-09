package main

import (
	"net/http"

	"github.com/fox-wei/go_project_training/life_revive/health"
	"github.com/gin-gonic/gin"
)

func Load(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "API not found!")
	})

	check := engine.Group("/check")
	{
		check.GET("/health", health.Health)
	}

	db := engine.Group("/db")
	{
		//*首页相关
		//?首页导航
		db.GET("/nav", NavHandler.NavHandler)
		db.GET("/subnav", NavHandler.SubNavHandler)

		//?获取团购信息
		db.GET("/team/detail/:id", TeamDetailHandler.TeamDetailHandler)

		//?我的页面配置
		db.GET("/me", MeHandler.MeHandler)
	}

	return engine
}
