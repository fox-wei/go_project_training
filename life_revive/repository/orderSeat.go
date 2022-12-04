package repository

import (
	"log"

	"github.com/fox-wei/go_project_training/life_revive/model"
)

type OrderSeatRepo struct {
	DB model.DataBase
}

//?订单查询
func (osr *OrderSeatRepo) OrderSeatOp(o model.OrderSeat) string {
	//?update
	err := osr.DB.MyDB.Save(o).Error

	if err != nil {
		log.Println(err)
	}

	return o.OrderSeatId
}
