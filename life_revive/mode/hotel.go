package mode

type Hotel struct {
	HotelId       string  `json:"hotelId"`
	HotelName     string  `json:"hotelName"`
	Pic           string  `json:"pic"`           //*图片地址
	Star          string  `json:"star"`          //*评分
	Num           int     `json:"num"`           //?评论数
	AveragePrice  string  `json:"averagePrice"`  //^人均价格
	Taste         float64 `json:"taste"`         //&口味
	Env           float64 `json:"env"`           //&环境
	Service       float64 `json:"service"`       //&服务
	MapAddr       string  `json:"mapaddr"`       //~ 详细地址(某陆某号广场2层)
	MapAddr2      string  `json:"mapaddr2"`      //~ 详细地址
	ShortType     string  `json:"shortType"`     //?类型
	BusinnessTime string  `json:"businessTimem"` //?营业时间
	Bang          string  `json:"bang"`          //?榜单排行

	TeamList       []Team        //*团购列表
	FoodList       []SuggestFood //*推荐菜单
	CommentTagList []CommentTag  //*评论tag表
	Comment        []Comment     //*评论表
	Market         Market        //*商场
}
