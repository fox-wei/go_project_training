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

		//?首页拼团
		db.GET("/team", SuggestFoodHandler.TeamHandler)
		db.GET("/rush", SuggestFoodHandler.RushHandler)

		//?首页猜你喜欢
		db.GET("/guess", GuessHandler.Guess)

		//*查找优惠信息
		db.GET("/discount", DiscountHandler.DiscountList)

		//?我的页面配置
		db.GET("/me", MeHandler.MeHandler)
	}

	return engine
}
