package model

//*团购套餐
type TeamChooseItem struct {
	TeamChooseItemId   string `gorm:"column:team_choose_item_id" json:"chooseItemId"`
	TeamChooseItemName string `gorm:"column:team_choose_item_name" json:"chooseItemName"` //?名称
	TeamChooseItemTip  string `gorm:"team_choose_item_tip" json:"chooseItemTip"`          //?提示如选取2份
	TeamPrice          int    `gorm:"column:team_price" json:"teamPrice"`                 //?价格
}
