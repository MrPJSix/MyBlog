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

type IArticleController interface {
	CreateArticle(c *gin.Context)
	GetArticleInfo(c *gin.Context)
	GetArticleList(c *gin.Context)
	GetListByTitle(c *gin.Context)
	GetListByCategory(c *gin.Context)
	GetListByUser(c *gin.Context)
	UpdateArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)
}

type ArticleController struct {
	articleService *service.ArticleService
}

func NewArticleController() *ArticleController {
	articleService := service.NewArticleService()
	return &ArticleController{articleService}
}

/* ====================================== */

func (ac *ArticleController) CreateArticle(c *gin.Context) {
	var data model.Article
	var code int
	err := c.ShouldBindJSON(&data)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		code = ac.articleService.CreateArticle(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, code := ac.articleService.GetArticleInfo(uint(id))
	responseData := dto.ArticleToResponse(&article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetArticleList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	var articles []model.Article
	var total int64
	var code int

	if len(title) == 0 {
		articles, total, code = ac.articleService.GetArticleList(pageSize, pageNum)
	} else {
		articles, total, code = ac.articleService.GetListByTitle(title, pageSize, pageNum)
	}

	responseData := dto.ArticleSliceToResponse(articles)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"total":   total,
		"meesage": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetListByCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	articles, total, code := ac.articleService.GetListByCategory(uint(id), pageSize, pageNum)

	responseData := dto.ArticleSliceToResponse(articles)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetListByUser(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	articles, total, code := ac.articleService.GetListByUser(uint(id), pageSize, pageNum)

	responseData := dto.ArticleSliceToResponse(articles)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) UpdateArticle(c *gin.Context) {
	var data model.Article
	var code int
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&data)

	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}

	code = ac.articleService.UpdateArticle(uint(id), &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := ac.articleService.DeleteArticle(uint(id))

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
