package service

import (
	"myblog.backend/model"
	"myblog.backend/repository"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type ICommentService interface {
	checkUserRight(requester *model.User, commenterID uint) bool
	CreateComment(comment *model.Comment) int
	GetCommentsByArticleID(articleID uint) ([]model.Comment, int64, int)
	DeleteComment(requester *model.User, id uint) int
}

type CommentService struct {
	commentRepo *repository.CommentRepo
}

func NewCommentService() *CommentService {
	commentRepo := repository.NewCommentRepo()
	return &CommentService{commentRepo}
}

/* ====================================== */

func (cs *CommentService) checkUserRight(requester *model.User, commenterID uint) bool {
	if requester.Role == 1 {
		return true
	}
	if requester.ID == commenterID {
		return true
	}
	return false
}

// 新增评论
func (cs *CommentService) CreateComment(comment *model.Comment) int {
	return cs.commentRepo.Create(comment)
}

// 获取某篇文章的所有评论
func (cs *CommentService) GetCommentsByArticleID(articleID uint) ([]model.Comment, int64, int) {
	return cs.commentRepo.GetByArticleID(articleID)
}

// 删除评论
func (cs *CommentService) DeleteComment(requester *model.User, id uint) int {
	comment, code := cs.commentRepo.GetByID(id)
	if code != errmsg.SUCCESS {
		return code
	}
	if !cs.checkUserRight(requester, comment.UserID) {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return cs.commentRepo.Delete(id)
}
