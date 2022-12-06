package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
)

type RestaurantService struct {
	Repo     *repository.RestauranNavRepo
	ItemRepo *repository.RestaurantTableRepo
}

func (r *RestaurantService) ListRestaurantNav(level int) []model.RestaurantNav {
	items := r.Repo.ListRestaurantNav(level)
	return items
}

func (r *RestaurantService) ListGoodRestaurantTableItem() []model.RestaurantTabItem {
	items := r.ItemRepo.ListGoodRestaurantTableItem()
	return items
}
