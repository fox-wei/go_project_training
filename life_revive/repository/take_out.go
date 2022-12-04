package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type TakeOutItemRepo struct {
	DB model.DataBase
}

func (t *TakeOutItemRepo) GetTakeOutById(hotelId string) []model.TakeOut {
	var itemlists []model.TakeOut

	t.DB.MyDB.Table("hotel").Joins("JOIN hotel_food_category on hotel.hotel_id=hotel_food_category.hotel_id").
		Joins("JOIN take_out on hotel_food_category.hotel_food_category_id=take_out.hotel_food_category_id").Where("hotel.hotel_id=?", hotelId).
		Select("hotel.hotel_name,hotel_food_category.*,take_out.*").Find(&itemlists)

	return itemlists
}
