package dbservice

import (
	"github.com/fox-wei/go_project_training/life_revive/model"
	"github.com/fox-wei/go_project_training/life_revive/repository"
)

type CommentTagService struct {
	Repo repository.CommentTagRepo
}

func (c *CommentTagService) GetCommentTagList(hoteId string) []model.CommentTag {
	return c.Repo.GetCommentTagList(hoteId)
}
