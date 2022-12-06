package dbservice

import (
	"errors"

	"github.com/fox-wei/go_project_training/life_revive/handler/param"
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
	"github.com/hashicorp/go-uuid"
)

type PostTeamOrderService struct {
	TeamRepo *repository.TeamRepo
	Repo     *repository.TeamPostOrderRepo
}

func (t *PostTeamOrderService) PostTeamOrder(order param.TeamPostOrder) (string, error) {
	//*是否存在这个团队优惠
	teamDetail := t.TeamRepo.GetTeamDetail(order.TeamDetailId)
	if teamDetail == nil {
		return "", errors.New("teamDetailId argument is error!")
	}

	//?下单数量不能小于1
	if order.Quantity < 1 {
		return "", errors.New("quantity argument is error!")
	}

	//?价格大于0
	if order.RealPrice < 0 {
		return "", errors.New("realPrice argument is error!")
	}

	//*手机号不能空
	if order.Message == "" {
		return "", errors.New("mobile can't empty!")
	}

	id, _ := uuid.GenerateUUID()
	o := model.TeamPostOrder{
		TeamPostId:   id,
		TeamDetailId: order.TeamDetailId,
		RealPrice:    order.RealPrice,
		Quantity:     order.Quantity,
		Mobile:       order.Mobile,
		Name:         order.Name,
		Sex:          order.Sex,
		Message:      order.Message,
	}

	return t.Repo.Save(o), nil
}
