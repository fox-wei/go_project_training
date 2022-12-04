package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type SuggestFoodRepo struct {
	DB model.DataBase
}

func (s *SuggestFoodRepo) GetFoodById(hotelId string) []model.SuggestFood {
	var foods []model.SuggestFood

	s.DB.MyDB.Where("hotel_i=?", hotelId).Find(&foods)
	return foods
}

func (s *SuggestFoodRepo) ListSuggest(level int) []model.Suggest {
	var items []model.Suggest

	s.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}
