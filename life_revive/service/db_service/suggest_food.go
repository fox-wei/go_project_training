package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
)

type SuggestFoodService struct {
	Repo *repository.SuggestFoodRepo
}

func (s *SuggestFoodService) ListSuggests(level int) []model.Suggest {
	return s.Repo.ListSuggest(level)
}

func (s *SuggestFoodService) GetFoodById(hotelId string) []model.SuggestFood {
	return s.Repo.GetFoodById(hotelId)
}

type SuggestService struct {
	Repo *repository.SuggestRepo
}

func (s *SuggestService) GetSuggestByLevel(level int) model.Suggest {
	return s.Repo.ListSuggest(level)
}
