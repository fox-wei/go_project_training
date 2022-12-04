package repository

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
)

type RestauranNavtRepo struct {
	DB model.DataBase
}

func (r *RestauranNavtRepo) ListRestaurantNav(level int) []model.RestaurantNav {
	var items []model.RestaurantNav

	r.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}

type RestaurantTableRepo struct {
	DB model.DataBase
}

func (r *RestaurantTableRepo) ListGoodRestaurantTable(level int) []model.RestaurantTabItem {
	var items []model.RestaurantTabItem

	r.DB.MyDB.Find(&items)
	return items
}

func (r *RestaurantTableRepo) ListRestaurantNav(level int) []model.RestaurantNav {
	var items []model.RestaurantNav

	r.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}
