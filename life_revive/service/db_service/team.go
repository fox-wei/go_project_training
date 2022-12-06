package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
	"github.com/fox-wei/go_project_training/life_revive/res"
)

type TeamService struct {
	Repo *repository.TeamRepo
}

func (t *TeamService) GetTeamListById(hotelId string) []model.Team {
	return t.Repo.GetTeamListById(hotelId)
}

func (t *TeamService) GetTeamDetail(teamId string) res.TeamResp {
	detail := t.Repo.GetTeamDetail(teamId)
	item := convert2TeamResp(detail)
	return item
}

func convert2TeamResp(items []model.TeamAggregation) res.TeamResp {
	var resp res.TeamResp
	resp.TeamDetail = items[0].TeamDetail
	for _, item := range items {
		f := &item.TeamChooseFood
		resp.FoodList = append(resp.FoodList, Convert2TeamChooseFoodResp(f))

		i := &item.TeamChooseItem
		resp.ChooseItems = append(resp.ChooseItems, TeamChooseItem(i))
	}

	return resp
}

func TeamChooseItem(item *model.TeamChooseItem) res.TeamChooseItemResp {
	r := res.TeamChooseItemResp{
		TeamChoseItemId:    item.TeamChooseItemId,
		TeamChooseItemName: item.TeamChooseItemName,
		TeamChooseItemTip:  item.TeamChooseItemTip,
		TeamPrice:          item.TeamPrice,
	}
	return r
}

func Convert2TeamChooseFoodResp(food *model.TeamChooseFood) res.TeamChooseFoodResp {
	return res.TeamChooseFoodResp{
		TeamChooseFoodId:   food.TeamChooseFoodId,
		TeamChooseFoodName: food.TeamChooseFoodName,
	}
}
