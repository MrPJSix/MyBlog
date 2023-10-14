package service

import (
	"myblog.backend/model"
	"myblog.backend/repository"
	"myblog.backend/utils/errmsg"
	"sort"
)

/* ====================================== */

type ICommentService interface {
	checkUserRight(requester *model.User, commenterID uint) bool
	CreateComment(comment *model.Comment) int
	GetCommentsByArticleID(articleID uint) ([]model.Comment, int64, int)
	DeleteComment(requester *model.User, id uint) int
	buildCommentReplies(uint, map[uint][]model.Comment) []model.Comment
	GetRootCommentsByArticleID(articleID uint) ([]model.Comment, int)
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

func (cs *CommentService) buildCommentReplies(rootID uint, childCommentsMap map[uint][]model.Comment) []model.Comment {
	replies := childCommentsMap[rootID]
	sort.Slice(replies, func(i, j int) bool {
		return replies[i].CreatedAt.Before(replies[j].CreatedAt)
	})
	return replies
}

func (cs *CommentService) GetRootCommentsByArticleID(articleID uint) ([]model.Comment, int) {
	var code int
	var rootComments []model.Comment
	rootComments, code = cs.commentRepo.GetRootByArticleID(articleID)
	if code != errmsg.SUCCESS {
		return nil, code
	}
	var allReplies []model.Comment
	allReplies, code = cs.commentRepo.GetRepliesByArticleID(articleID)
	if code != errmsg.SUCCESS {
		return nil, code
	}
	childCommentsMap := make(map[uint][]model.Comment)
	for _, reply := range allReplies {
		childCommentsMap[reply.RootCommentID] = append(childCommentsMap[reply.RootCommentID], reply)
	}

	// 为每个根评论构建评论树
	for idx, comment := range rootComments {
		rootComments[idx].Replies = cs.buildCommentReplies(comment.ID, childCommentsMap)
	}

	return rootComments, errmsg.SUCCESS
}
