package model

//?团购套餐
type TeamChooseFood struct {
	TeamChooseFoodId   string `gorm:"column:team_choose_food_id" json:"chooseFoodId"`
	TeamChooseFoodName string `gorm:"column:team_choose_food_Name" json:"chooseName"`
}
