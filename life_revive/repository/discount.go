package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type Discount struct {
	DB model.DataBase
}

func (d *Discount) ListDiscounts(pos int) []model.Discount {
	var discounts []model.Discount

	d.DB.MyDB.Where("pos=?", pos).Find(&discounts)
	return discounts
}
