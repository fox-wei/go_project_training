package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type CommentTagRepo struct {
	DB model.DataBase
}

func (c *CommentTagRepo) GetCommentTagList(hotelId string) []model.CommentTag {
	var tagList []model.CommentTag
	c.DB.MyDB.Where("hotel_id=?", hotelId).Find(&tagList)
	return tagList
}
