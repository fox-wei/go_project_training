package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fox-wei/go_project_training/life_revive/handler/param"
	dbservice "github.com/fox-wei/go_project_training/life_revive/service/db_service"
	"github.com/gin-gonic/gin"
)

const domain = "http://localhost:8080"

type DiscountHandler struct {
	Srv *dbservice.DiscountService
}

type RestaurantNavHandler struct {
	Srv *dbservice.RestaurantService
}

type HotelDetailHandler struct {
	Srv *dbservice.HotelService
}

type TeamDetailHandler struct {
	Srv *dbservice.TeamService
}

type OrderSeatHandler struct {
	Srv *dbservice.OrderSeatService
}

type TakeOutHandler struct {
	Srv *dbservice.TakeOutService
}

type MeHandler struct {
	Srv *dbservice.MeService
}

type SuggestFoodHandler struct {
	Srv *dbservice.SuggestFoodService
}

type SuggestHandler struct {
	Srv *dbservice.SuggestService
}

type GuessHandler struct {
	Srv *dbservice.GuessService
}

type NavHandler struct {
	Srv *dbservice.ListNavItemService
}

type PostTeamOrderHandler struct {
	Srv *dbservice.PostTeamOrderService
}

//*逻辑实现
func (h *DiscountHandler) DiscountList(ctx *gin.Context) {
	discounts := h.Srv.ListDiscounts()
	ctx.JSON(http.StatusOK, gin.H{
		"items": discounts,
	})
}

func (h *RestaurantNavHandler) RestaurantNav(ctx *gin.Context) {
	items := h.Srv.ListRestaurantNav(1) //?[]model.RestaurantNav
	ctx.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *RestaurantNavHandler) GoodRestaurantBillBoardHandler(ctx *gin.Context) {
	items := h.Srv.ListRestaurantNav(2)
	ctx.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *RestaurantNavHandler) GoodRestaurantTableItemHandler(ctx *gin.Context) {
	items := h.Srv.ListGoodRestaurantTableItem() //?[]model.RestaurantTableItem
	ctx.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *HotelDetailHandler) HotelDetailHandler(ctx *gin.Context) {
	hotelId := ctx.Param("id") //?路由参数获取

	hotel, err := h.Srv.GetHotelDetailById(hotelId)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"item": nil,
			"msg":  err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"item": hotel,
			"msg":  "",
		})
	}
}

func (h *TeamDetailHandler) TeamDetailHandler(ctx *gin.Context) {
	teamId := ctx.Param("id")
	detail := h.Srv.GetTeamDetail(teamId)

	ctx.JSON(http.StatusOK, gin.H{
		"item": detail,
	})
}

func (h *PostTeamOrderHandler) TeamOrderHandler(ctx *gin.Context) {
	var p param.TeamPostOrder

	err := ctx.BindJSON(&p) //?绑定，接收订单信息

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"id":  "",
			"err": err.Error(),
		})
	}

	id, err := h.Srv.PostTeamOrder(p)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"id":  "",
			"err": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"id":  id,
			"err": "",
		})
	}

}

func (h *OrderSeatHandler) OrderSeat(ctx *gin.Context) {
	var p param.OrderSeat
	err := ctx.BindJSON(&p)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}

	id := h.Srv.OrderSeatRepoOp(p)
	ctx.JSON(http.StatusOK, gin.H{
		"id":  id,
		"err": "",
	})
}

func (h *TakeOutHandler) GetTakeOutById(ctx *gin.Context) {
	hotelId := ctx.Param("hotelId")
	item := h.Srv.GetTakeOutListById(hotelId)
	ctx.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func (h *MeHandler) MeHandler(ctx *gin.Context) {
	items := h.Srv.ListMe()
	ctx.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *NavHandler) NavHandler(ctx *gin.Context) {
	items := h.Srv.ListNavItems(1)
	ctx.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *NavHandler) SubNavHandler(ctx *gin.Context) {
	items := h.Srv.ListNavItems(2)
	ctx.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *SuggestFoodHandler) TeamHandler(ctx *gin.Context) {
	items := h.Srv.ListSuggests(1)
	ctx.JSON(http.StatusOK, gin.H{
		"item": items,
	})
}

func (h *SuggestFoodHandler) RushHandler(ctx *gin.Context) {
	items := h.Srv.ListSuggests(2)
	ctx.JSON(http.StatusOK, gin.H{
		"item": items,
	})
}

func (h *SuggestHandler) TeamHandler(ctx *gin.Context) {
	fmt.Println("TeamHandler...")
	items := h.Srv.GetSuggestByLevel(1)
	fmt.Println(items)
	ctx.JSON(http.StatusOK, gin.H{
		"item": items,
	})
}

func (h *SuggestHandler) RushHandler(ctx *gin.Context) {
	items := h.Srv.GetSuggestByLevel(2)
	ctx.JSON(http.StatusOK, gin.H{
		"item": items,
	})
}

func (h *GuessHandler) Guess(ctx *gin.Context) {
	items := h.Srv.ListGuesses()
	ctx.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func ImageHandler(ctx *gin.Context) {
	imageName := ctx.Query("imageName")
	fmt.Println(imageName)
	dir, _ := os.Getwd()
	file, ok := ioutil.ReadFile(dir + "/static/images/" + imageName + ".png")
	if ok != nil {
		fmt.Println(ok.Error())
	}
	ctx.Writer.WriteString(string(file))
}
