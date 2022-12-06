package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
)

type GuessService struct {
	Repo *repository.GuessRepo
}

func (g *GuessService) ListGuesses() []model.Guess {
	return g.Repo.ListGuesses()
}
