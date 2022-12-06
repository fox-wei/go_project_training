package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
)

type CommentService struct {
	Repo repository.CommentRepo
}

func (c *CommentService) GetCommentLst(hotelId string) []model.Comment {
	return c.Repo.GetCommentList(hotelId)
}
