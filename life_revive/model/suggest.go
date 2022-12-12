package model

//*首页推荐
type Suggest struct {
	SuggestId string `json:"suggestId"`

	//?keyword
	Keyword1 string `json:"keyword1"`
	Keyword2 string `json:"keyword2"`
	Keyword3 string `json:"keyword3"`
	Keyword4 string `json:"keyword4"`

	Src string `json:"src"`

	FoodName string `json:"foodName"`
	Icon     string `json:"icon"` //?icon图标

	HotelName   string `json:"hotelName"`
	JoinPersons string `json:"joinPersons"` //?参加人数
	Price       string `json:"price"`
	GoodPrice   string `json:"goodPrice"` //?优惠价格
	Distance    string `json:"distance"`  //?距离
	Level       int    `json:"level"`     //?等级分类
}
