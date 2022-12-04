package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type HotelRepo struct {
	DB model.DataBase
}

func (h *HotelRepo) GetHotelById(hotelId string) model.Hotel {
	var hotel model.Hotel

	//?select * from hotel where hotel_id=hotelId
	h.DB.MyDB.Where("hotel_id=?", hotelId).Find(&hotel)
	return hotel
}
