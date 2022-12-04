package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type TeamRepo struct {
	DB model.DataBase
}

func (t *TeamRepo) GetTeamListById(hotelId string) []model.Team {
	var team []model.Team

	t.DB.MyDB.Where("hotel_id=?", hotelId).Find(&team)
	return team
}

func (t *TeamRepo) GetTeamDetail(teamDetailId string) []model.TeamAggregation {
	var teamAggregations []model.TeamAggregation

	t.DB.MyDB.Table("team_detail").Where("team_detail.team_detail_id = ?", teamDetailId).
		Joins("JOIN team_choose_food on team_detail.team_detail_id=team_choose_food.team_detail_id").
		Joins("JOIN team_choose_item on team_choose_food.team_choose_food_id=team_choose_item.team_choose_food_id").
		Select("team_detail.*,team_choose_food.*,team_choose_item.*").Scan(&teamAggregations)
	return teamAggregations
}
