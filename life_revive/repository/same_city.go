package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type SameCityRepo struct {
	DB model.DataBase
}

func (s *SameCityRepo) ListSameCity() []model.SameCity {
	var sameCities []model.SameCity

	s.DB.MyDB.Find(&sameCities)
	return sameCities
}
