package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
)

type DiscountService struct {
	Repo *repository.Discount
}

func (d *DiscountService) ListDiscounts() []model.Discount {
	itemL := d.Repo.ListDiscounts(1)
	itemR := d.Repo.ListDiscounts(2)
	itemL = append(itemL, itemR...)
	return itemL
}
