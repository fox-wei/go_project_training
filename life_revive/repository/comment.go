package repository

import "github.com/fox-wei/go_project_training/life_revive/model"

type CommentRepo struct {
	DB model.DataBase
}

func (c *CommentRepo) GetCommentList(hoteId string) []model.Comment {
	var commentList []model.Comment
	c.DB.MyDB.Where("hotel_id=?", hoteId).Find(&commentList)
	return commentList
}
