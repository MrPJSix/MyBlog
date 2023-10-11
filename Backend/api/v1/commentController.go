package v1

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/dto"
	"myblog.backend/model"
	"myblog.backend/service"
	"myblog.backend/utils/errmsg"
	"net/http"
	"strconv"
)

/* ====================================== */

type ICommentController interface {
	CreateComment(c *gin.Context)
	GetCommentsByArticleID(c *gin.Context)
	DeleteComment(c *gin.Context)
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
func (cc *CommentController) CreateComment(c *gin.Context) {
	var comment model.Comment
	var code int

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		comment.UserID = c.MustGet("user_id").(uint)
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

// 获取某篇文章的所有评论
func (cc *CommentController) GetCommentsByArticleID(c *gin.Context) {
	aid, _ := strconv.Atoi(c.Param("id"))
	articleID := uint(aid)

	comments, total, code := cc.commentService.GetCommentsByArticleID(articleID)
	responseData := dto.CommentSliceToResponse(comments)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"total":   total,
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
