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

type ICommentController interface {
	CreateComment(comment *model.Comment) int
	GetCommentsByArticleID(articleID uint) ([]model.Comment, int64, int)
	DeleteComment(id uint) int
}

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController() *CommentController {
	commentService := service.NewCommentService()
	return &CommentController{commentService}
}

// 新增评论
func (cc *CommentController) CreateComment(c *gin.Context) {
	var comment model.Comment
	var code int

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		code = cc.commentService.CreateComment(&comment)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
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

	code := cc.commentService.DeleteComment(commentID)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
