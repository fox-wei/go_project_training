package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type GuessRepo struct {
	DB model.DataBase
}

func (g *GuessRepo) ListGuesses() []model.Guess {
	var guesses []model.Guess

	g.DB.MyDB.Find(&guesses)
	return guesses
}
