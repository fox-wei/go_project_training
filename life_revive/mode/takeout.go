package mode

//&外卖

type TakeOut struct {
	TakeOutId           string `json:"takeOutId"`
	HotelFoodCategoryId string `json:"hotelFoodCategoryId"` //&餐馆分类id

	FoodName      string  `json:"foodName"`
	Pic           string  `json:"pic"`          //*菜品图片
	MonthSoldNum  string  `json:"monthSoldNum"` //*月销量
	Zan           int     `json:"zan"`          //*好评率
	Price         int     `json:"price"`
	IsSuggest     int     `json:"isSuggest"`     //?是否推荐
	DiscountPrice int     `json:"discountPrice"` //?折扣价格
	Discount      float64 `json:"discount"`
}

type TakeOutItem struct {
	Hotel
	HotelFoodCategory
	TakeOut
}
