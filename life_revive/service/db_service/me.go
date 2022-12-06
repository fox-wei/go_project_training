package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
)

type MeService struct {
	Repo repository.MeItemRepo
}

func (m *MeService) ListMe() []model.MeItem {
	return m.Repo.ListMe()
}
