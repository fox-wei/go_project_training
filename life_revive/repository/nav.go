package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type ListNavItemRepo struct {
	DB model.DataBase
}

func (l *ListNavItemRepo) ListNavItems(level int) []model.Nav {
	var items []model.Nav

	l.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}
