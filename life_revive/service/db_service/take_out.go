package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
	"github.com/fox-wei/go_project_training/life_revive/res"
)

type TakeOutService struct {
	Repo *repository.TakeOutItemRepo
}

func Convert2HotelFoodCategoryResp(c model.HotelFoodCategory) res.HotelFoodCategoryResp {
	r := res.HotelFoodCategoryResp{
		HotelFoodCategoryId:   c.HotelFoodCategoryId,
		HotelFoodCategoryName: c.HotelFoodCategoryName,
		HotelId:               c.HotelId,
	}

	return r
}

func convert2TakeOutItemResp(t *model.TakeOut) res.TakeOutItemResp {
	r := res.TakeOutItemResp{
		TakeOutId:           t.TakeOutId,
		HotelFoodCategoryId: t.HotelFoodCategoryId,
		FoodName:            t.FoodName,
		Pic:                 t.Pic,
		MonthSoldNum:        t.MonthSoldNum,
		Zan:                 t.Zan,
		Price:               t.Price,
		IsSuggest:           t.IsSuggest,
		DiscountPrice:       t.DiscountPrice,
		Discount:            t.Discount,
	}

	return r
}

func (tos *TakeOutService) GetTakeOutListById(hotelId string) res.TakeOutResp {
	//!model.TakeOutItem
	items := tos.Repo.GetTakeOutListById(hotelId)
	var takeOutResp res.TakeOutResp

	if len(items) > 0 {
		takeOutResp.HotelName = items[0].HotelName
		for _, item := range items {
			c := model.HotelFoodCategory{
				HotelFoodCategoryId:   item.HotelFoodCategory.HotelFoodCategoryId,
				HotelFoodCategoryName: item.HotelFoodCategory.HotelFoodCategoryName,
				HotelId:               item.HotelFoodCategory.HotelId,
			}

			takeOutResp.CategoryList = append(takeOutResp.CategoryList, Convert2HotelFoodCategoryResp(c))

			t := model.TakeOut{
				TakeOutId:           item.TakeOutId,
				HotelFoodCategoryId: item.HotelFoodCategory.HotelFoodCategoryId,
				FoodName:            item.FoodName,
				Pic:                 item.TakeOut.Pic,
				MonthSoldNum:        item.MonthSoldNum,
				Zan:                 item.Zan,
				Price:               item.Price,
				IsSuggest:           item.IsSuggest,
				DiscountPrice:       item.DiscountPrice,
				Discount:            item.Discount,
			}
			takeOutResp.ItemList = append(takeOutResp.ItemList, convert2TakeOutItemResp(&t))
		}
	}
	return takeOutResp
}
