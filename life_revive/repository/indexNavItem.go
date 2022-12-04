package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type IndexNavItemRepo struct {
	DB model.DataBase
}

func (item *IndexNavItemRepo) ListIndexNavItem() ([]*model.IndexNavItem, uint64, error) {
	navItems := make([]*model.IndexNavItem, 0)
	var count uint64

	//*select count(*) from index_nav_item
	//?统计首页导航标签数量
	if err := item.DB.MyDB.Model(navItems).Count(count).Error; err != nil {
		return nil, 0, err
	}

	return navItems, count, nil
}
