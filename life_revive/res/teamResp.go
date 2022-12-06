package res

import "github.com/fox-wei/go_project_training/life_revive/model"

type TeamResp struct {
	model.TeamDetail
	FoodList    []TeamChooseFoodResp
	ChooseItems []TeamChooseItemResp
}
