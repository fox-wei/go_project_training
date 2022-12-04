package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type SuggestRepo struct {
	DB model.DataBase
}

func (s *SuggestRepo) ListSuggest(level int) model.Suggest {
	var suggest model.Suggest

	s.DB.MyDB.Where("level=?", level).Find(&suggest)
	return suggest
}
