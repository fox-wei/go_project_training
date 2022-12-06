package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
)

type MarketService struct {
	Repo repository.MarketRepo
}

func (m *MarketService) GetMarketInfo(hotelId string) model.Market {
	// var market model.Market
	// m.Repo.DB.MyDB.Where("hotel_id=?", hotelId).Find(&market)
	market := m.Repo.GetMarketInfo(hotelId)
	return market
}
