package model

type HotelFoodCategory struct {
	HotelFoodCategoryId   string `json:"hotelFoodCategoryId"`
	HotelFoodCategoryName string `json:"hotelFoodCategoryName"` //?分类名称
	HotelId               string `json:"hotelId"`               //*餐馆id
}
