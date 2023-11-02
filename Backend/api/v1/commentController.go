package v1

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/dto/response"
	"myblog.backend/middleware/auth"
	"myblog.backend/model"
	"myblog.backend/service"
	"myblog.backend/utils/errmsg"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/* ====================================== */

type ICommentController interface {
	CreateCommentToArticle(c *gin.Context)
	GetCommentsByArticleID(c *gin.Context)
	DeleteComment(c *gin.Context)
	GetRootCommentsByArticleID(c *gin.Context)
	GetRepliesByRootComment(c *gin.Context)
	GetAllCommentsCount(c *gin.Context)
	UserIsLiked(c *gin.Context)
	UserLikesComment(c *gin.Context)
}

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController() *CommentController {
	commentService := service.NewCommentService()
	return &CommentController{commentService}
}

/* ====================================== */

// 新增评论
func (cc *CommentController) CreateCommentToArticle(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))
	var comment model.Comment
	var code int

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		comment.UserID = c.MustGet("user_id").(uint)
		comment.ArticleID = uint(articleID)
		code = cc.commentService.CreateComment(&comment)
	}

	var responseData *dto.CommentResponse
	if code == errmsg.SUCCESS {
		responseData = dto.CommentToResponse(&comment)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

// 新增回复
func (cc *CommentController) CreateReply(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))
	commentID, _ := strconv.Atoi(c.Param("cid"))
	parentCommentID := uint(commentID)

	var reply model.Comment
	var code int
	err := c.ShouldBindJSON(&reply)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		reply.UserID = c.MustGet("user_id").(uint)
		reply.ArticleID = uint(articleID)
		reply.ParentCommentID = &parentCommentID
		code = cc.commentService.CreateComment(&reply)
	}

	var responseData *dto.CommentResponse
	if code == errmsg.SUCCESS {
		responseData = dto.CommentToResponse(&reply)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

// 获取某篇文章的所有评论
func (cc *CommentController) GetCommentsByArticleID(c *gin.Context) {
	aid, _ := strconv.Atoi(c.Param("id"))
	articleID := uint(aid)

	comments, total, code := cc.commentService.GetCommentsByArticleID(articleID)

	tokenHeader := c.GetHeader("Authorization")
	checkToken := strings.SplitN(tokenHeader, " ", 2)
	var userID uint
	if tokenHeader == "" || (len(checkToken) != 2 && checkToken[0] != "Bearer") {
		userID = 0
	}
	claims, code2 := auth.CheckToken(checkToken[1])
	if code2 == errmsg.ERROR || time.Now().Unix() > (*claims.ExpiresAt).Time.Unix() {
		userID = 0
	} else {
		userID = claims.UserID
	}

	responseData := dto.CommentLikeSliceToResponse(comments, userID)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func (cc *CommentController) GetRootCommentsByArticleID(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))
	var comments []model.Comment
	var code int
	comments, code = cc.commentService.GetRootCommentsByArticleID(uint(articleID))

	tokenHeader := c.GetHeader("Authorization")
	checkToken := strings.SplitN(tokenHeader, " ", 2)
	var responseData []*dto.CommentResponse
	if tokenHeader == "" || (len(checkToken) != 2 && checkToken[0] != "Bearer") {
		responseData = dto.CommentSliceToResponse(comments)
	} else if claims, code2 := auth.CheckToken(checkToken[1]); code2 == errmsg.ERROR || time.Now().Unix() > (*claims.ExpiresAt).Time.Unix() {
		responseData = dto.CommentSliceToResponse(comments)
	} else {
		responseData = dto.CommentLikeSliceToResponse(comments, claims.UserID)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除评论
func (cc *CommentController) DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	commentID := uint(id)
	var requester model.User
	requester.ID = c.MustGet("user_id").(uint)
	requester.Role = c.MustGet("role").(uint8)

	code := cc.commentService.DeleteComment(&requester,
		commentID)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func (cc *CommentController) GetRepliesByRootComment(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	replies, code := cc.commentService.GetRepliesByRoot(uint(id), pageSize, pageNum)

	tokenHeader := c.GetHeader("Authorization")
	checkToken := strings.SplitN(tokenHeader, " ", 2)
	var responseData []*dto.CommentResponse
	if tokenHeader == "" || (len(checkToken) != 2 && checkToken[0] != "Bearer") {
		responseData = dto.CommentSliceToResponse(replies)
	} else if claims, code2 := auth.CheckToken(checkToken[1]); code2 == errmsg.ERROR || time.Now().Unix() > (*claims.ExpiresAt).Time.Unix() {
		responseData = dto.CommentSliceToResponse(replies)
	} else {
		responseData = dto.CommentLikeSliceToResponse(replies, claims.UserID)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})

}

func (cc *CommentController) GetAllCommentsCount(c *gin.Context) {
	total, code := cc.commentService.GetAllCommentsCount()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    total,
		"message": errmsg.GetErrMsg(code),
	})
}

func (cc *CommentController) UserIsLiked(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("user_id").(uint)

	data, code := cc.commentService.UserIsLiked(uint(id), userID)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func (cc *CommentController) UserLikesComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("user_id").(uint)

	code := cc.commentService.UserLikesComment(uint(id), userID)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
