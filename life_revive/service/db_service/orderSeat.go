package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/handler/param"
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
	"github.com/hashicorp/go-uuid"
)

type OrderSeatService struct {
	Repo *repository.OrderSeatRepo
}

func (oss *OrderSeatService) OrderSeatRepoOp(p param.OrderSeat) string {
	//*order.HotelId连接数据库查询，是否存在这个团购优惠
	//*防止恶意篡改优惠
	//*order.time是否超过选择范围
	//*order.mobile 格式是否正确
	//?rootype检查预订类型是否合法
	//?成功预订后，发送短信
	id, _ := uuid.GenerateUUID() //?生成uuid(通用唯一识别码)
	o := model.OrderSeat{
		OrderSeatId: id,
		HotelId:     p.HotelId,
		Mobile:      p.Mobile,
		Name:        p.Name,
		Sex:         p.Sex,
		Message:     p.Message,
	}

	return oss.Repo.OrderSeatOp(o)
}
