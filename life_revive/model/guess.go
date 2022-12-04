package model

//!guess what you like
type Guess struct {
	GuessId   string `json:"guess_id"`
	Src       string `json:"src"` //*picture of address
	Title     string `json:"title"`
	Desc      string `json:"desc"`      //*内容
	GoodPrice string `json:"goodPrice"` //&优惠价格
	Price     string `json:"price"`     //&原价
	SoldNum   int32  `json:"soldNum"`   //*售出分数
}
