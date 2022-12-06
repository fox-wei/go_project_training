package res

//*团购套餐选取
type TeamChooseItemResp struct {
	TeamChoseItemId    string `gorm:"column:team_chose_item_id" json:"choseItemId"`
	TeamChooseItemName string `gorm:"column:team_choose_item_name" json:"choseItemName"`
	TeamChooseItemTip  string `gorm:"column:team_choose_item_tip" json:"chooseItemTip"` //?提示
	TeamPrice          int    `gorm:"column:team_price" json:"teamPrice"`
}
