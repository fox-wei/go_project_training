package mode

type Discount struct {
	Id                    string `json:"id"` //^自增id
	DiscountId            string `json:"discountId"`
	DiscountItemIcon      string `json:"discountItemIcon"`      //*优惠icon地址
	DiscountItemSrc       string `json:"discountItemSrc"`       //*优惠项目src
	DiscountItemTitle     string `json:"discountItemTitle"`     //*优惠项目标题
	DiscountItemHotel     string `json:"discountItemHotel"`     //*优惠餐馆
	DiscountItemGoodPrice int32  `json:"discountItemGoodPrice"` //?优惠价格
	DiscountItemPrice     int32  `json:"discountItemPrice"`     //?原价格
	Discount              string `json:"discount"`              //^折扣
	Pos                   int    `json:"pos"`                   //&最会做
}
