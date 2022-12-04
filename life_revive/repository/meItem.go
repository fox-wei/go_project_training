package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type MeItemRepo struct {
	DB model.DataBase
}

func (me *MeItemRepo) ListMe() []model.MeItem {
	var items []model.MeItem

	me.DB.MyDB.Find(&items)
	return items
}
