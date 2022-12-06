package dbservice

import (
	"errors"

	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
)

type HotelService struct {
	Repo            *repository.HotelRepo
	TeamRepo        *repository.TeamRepo
	SuggestFoodRepo *repository.SuggestFoodRepo
	CommentTagRepo  *repository.CommentTagRepo
	CommentRepo     *repository.CommentRepo
	MarketRepo      *repository.MarketRepo
}

func (h *HotelService) GetHotelDetailById(id string) (*model.Hotel, error) {
	if id == "" {
		return nil, errors.New("参数错误！")
	}

	hotel := h.Repo.GetHotelById(id)
	if &hotel == nil { //?结构体
		return nil, errors.New("餐馆查询错误！")
	}

	teamList := h.TeamRepo.GetTeamListById(id)
	hotel.TeamList = teamList

	foodList := h.SuggestFoodRepo.GetFoodById(id)
	hotel.FoodList = foodList

	tagList := h.CommentTagRepo.GetCommentTagList(id)
	hotel.CommentTagList = tagList

	commentList := h.CommentRepo.GetCommentList(id)
	hotel.Comment = commentList

	market := h.MarketRepo.GetMarketInfo(id)
	hotel.Market = market
	return &hotel, nil

}
