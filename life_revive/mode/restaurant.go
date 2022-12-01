package mode

//*美食热门排行
type RestaurantNav struct {
	NavId string `json:"navId"`
	Title string `json:"title"`
	Desc  string `json:"desc"`  //?描述
	Src   string `json:"src"`   //?图片地址
	Level string `json:"level"` //?等级
}

type RestaurantTabItem struct {
	ItemId string  `json:"itemId"`
	Src    string  `json:"src"`
	Title  string  `json:"title"`
	Star   float32 `json:"star"`
	Price  string  `json:"price"`
	Area   string  `json:"area"`
	Type   string  `json:"type"`
	Desc   string  `json:"desc"`
	Team   string  `json:"team"`
	Quan   string  `json:"quan"`
}
