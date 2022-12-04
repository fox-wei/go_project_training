package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type TeamPostOrderRepo struct {
	DB model.DataBase
}

func (t *TeamPostOrderRepo) Save(order model.TeamPostOrder) string {
	tx := t.DB.MyDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(order).Error
	if err != nil {
		tx.Rollback()
		return ""
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return ""
	}

	return order.TeamDetailId
}
