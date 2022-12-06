package res

type TakeOutResp struct {
	HotelName    string `josn:"hotelName"`
	CategoryList []HotelFoodCategoryResp
	ItemList     []TakeOutItemResp
}
