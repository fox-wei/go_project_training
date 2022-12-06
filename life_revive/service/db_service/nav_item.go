package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
	"github.com/fox-wei/go_project_training/life_revive/res"
)

type ListNavItemService struct {
	Repo *repository.ListNavItemRepo
}

func model2res(nav model.Nav) res.Nav {
	navResp := res.Nav{
		NavId: nav.NavId,
		Src:   nav.Src,
		Title: nav.Title,
	}
	return navResp
}

func (l *ListNavItemService) ListNavItems(level int) []res.Nav {
	result := l.Repo.ListNavItems(level)

	var newList []res.Nav

	for _, item := range result {
		r := model2res(item)
		newList = append(newList, r)
	}

	return newList
}
