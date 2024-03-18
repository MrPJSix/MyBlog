package service

import (
	"log"
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
	GetAllCommentsCount() (int64, int)

	// 点赞功能
	UserIsLiked(commentID, userID uint) (bool, int)
	likeSQLToRedis(commentID uint)
	UserLikesComment(commentID, userID uint) int
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

// 获取根评论
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
		childCommentsMap[*reply.RootCommentID] = append(childCommentsMap[*reply.RootCommentID], reply)
	}

	// 为每个根评论构建评论树
	for idx, comment := range rootComments {
		rootComments[idx].Replies = cs.buildCommentReplies(comment.ID, childCommentsMap)
	}

	return rootComments, errmsg.SUCCESS
}

// 获取根评论的回复
func (cs *CommentService) GetRepliesByRoot(rootCommentID uint, pageSize, pageNum int) ([]model.Comment, int) {
	var offset int
	if pageSize >= 100 {
		pageSize = 100
	} else if pageSize <= 0 {
		pageSize = -1
	}
	if pageNum <= 0 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	relies, code := cs.commentRepo.GetRepliesByRoot(rootCommentID, pageSize, offset)
	if code != errmsg.SUCCESS {
		return nil, code
	}
	return relies, code
}

// 获取评论总数
func (cs *CommentService) GetAllCommentsCount() (int64, int) {
	return cs.commentRepo.GetAllCount()
}

// 点赞功能
func (cs *CommentService) UserIsLiked(commentID, userID uint) (bool, int) {
	var code int
	code = cs.commentRepo.UserIsLikedRds(commentID, userID)
	if code == errmsg.REDIS_SET_IS_MEMBER {
		return true, errmsg.SUCCESS
	} else if code == errmsg.REDIS_SET_ISNOT_MEMBER {
		return false, errmsg.SUCCESS
	} else if code == errmsg.REDIS_SET_NOT_EXISTS {
		go cs.likeSQLToRedis(commentID)
	} else if code == errmsg.REDIS_IS_SYNCING {
	}
	return cs.commentRepo.UserIsLikedSQL(commentID, userID)
}

// 用Redis太复杂
func (cs *CommentService) likeSQLToRedis(commentID uint) {
	err := cs.commentRepo.SaveLikesToRedis(commentID)
	if err != nil {
		log.Println("评论点赞加载到Redis出错", commentID, err)
	} else {
		log.Println("评论点赞加载到Redis成功", commentID)
	}
}

func (cs *CommentService) UserLikesComment(commentID, userID uint) int {
	var code int
	rdsCode := cs.commentRepo.UserIsLikedRds(commentID, userID)
	if rdsCode == errmsg.REDIS_SET_IS_MEMBER {
		go cs.commentRepo.DecreaseLikes(commentID, userID)
		return cs.commentRepo.DecreaseLikesRds(commentID, userID)
	} else if rdsCode == errmsg.REDIS_SET_ISNOT_MEMBER {
		go cs.commentRepo.IncreaseLikes(commentID, userID)
		return cs.commentRepo.IncreaseLikesRds(commentID, userID)
	} else if rdsCode == errmsg.REDIS_IS_SYNCING {
		isLiked, code := cs.commentRepo.UserIsLikedSQL(commentID, userID)
		if code != errmsg.SUCCESS {
			return code
		}
		if isLiked {
			go cs.commentRepo.DecreaseLikesRds(commentID, userID)
			code = cs.commentRepo.DecreaseLikes(commentID, userID)
		} else {
			go cs.commentRepo.IncreaseLikesRds(commentID, userID)
			code = cs.commentRepo.IncreaseLikes(commentID, userID)
		}
		return code
	}

	isLiked, code := cs.UserIsLiked(commentID, userID)
	if code != errmsg.SUCCESS {
		return code
	}
	if isLiked {
		code = cs.commentRepo.DecreaseLikes(commentID, userID)
	} else {
		code = cs.commentRepo.IncreaseLikes(commentID, userID)
	}

	return code
}
