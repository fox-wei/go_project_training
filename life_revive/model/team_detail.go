package model

type TeamDetail struct {
	TeamDetailId   string `gorm:"column:team_detail_id" json:"teamDetailId"`
	TeamDetailName string `gorm:"column:team_detail_id" json:"teamDetailName"` //*团购详细名称
	Policy         string `gorm:"column:policy" json:"policy"`                 //*团购政策
	Tips           string `gorm:"column:tips" json:"tips"`
}

type TeamAggregation struct {
	TeamDetail
	TeamChooseFood
	TeamChooseItem
}
