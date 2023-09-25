package service

import (
	"myblog.backend/model"
	"myblog.backend/repository"
)

/* ====================================== */

type ICommentService interface {
	CreateComment(comment *model.Comment) int
	GetCommentsByArticleID(articleID uint) ([]model.Comment, int64, int)
	DeleteComment(id uint) int
}

type CommentService struct {
	commentRepo *repository.CommentRepo
}

func NewCommentService() *CommentService {
	commentRepo := repository.NewCommentRepo()
	return &CommentService{commentRepo}
}

/* ====================================== */

// 新增评论
func (cs *CommentService) CreateComment(comment *model.Comment) int {
	return cs.commentRepo.Create(comment)
}

// 获取某篇文章的所有评论
func (cs *CommentService) GetCommentsByArticleID(articleID uint) ([]model.Comment, int64, int) {
	return cs.commentRepo.GetByArticleID(articleID)
}

// 删除评论
func (cs *CommentService) DeleteComment(id uint) int {
	return cs.commentRepo.Delete(id)
}
