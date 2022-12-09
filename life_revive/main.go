package main

import (
	"flag"
	"log"
	"net/http"

	mylog "github.com/fox-wei/go_project_training/life_revive/MyLog"
	"github.com/fox-wei/go_project_training/life_revive/config"
	"github.com/fox-wei/go_project_training/life_revive/handler"
	"github.com/fox-wei/go_project_training/life_revive/middleware"
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	cfg                  = flag.String("config", "", "")
	DiscountHandler      handler.DiscountHandler
	RestaurantNavHandler handler.RestaurantNavHandler
	HotelDetailHandler   handler.HotelDetailHandler
	TeamDetailHandler    handler.TeamDetailHandler
	OrderSeatHandler     handler.OrderSeatHandler
	TakeOutHandler       handler.TakeOutHandler
	MeHandler            handler.MeHandler
	SuggestFoodHandler   handler.SuggestFoodHandler
	SuggestHandler       handler.SuggestHandler
	GuessHandler         handler.GuessHandler
	NavHandler           handler.NavHandler
	PostTeamOrderHandler handler.PostTeamOrderHandler
)

func init() {
	initViper()
	initDB()
	initHandler()
}

func main() {

	//*解析配置文件
	flag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))

	model.DB.Init() //*初始化DB
	defer model.DB.Close()

	r := gin.New()

	Load(r, middleware.ProcessTraceID(), middleware.Logging())

	port := viper.GetString("addr")
	mylog.Log.Info("监听端口", port)
	log.Println(http.ListenAndServe(port, r).Error())

}
