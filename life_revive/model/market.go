package model

type Market struct {
	MarketId   string `json:"marketId"`
	MarketName string `json:"marketName"`
	Addr       string `json:"addr"`
	ShortType  string `json:"shortType"` //&类型
	Star       int    `json:"star"`
	Pic        string `json:"pic"` //&图片
}
