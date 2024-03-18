package v1

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/dto/response"
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
	GetAllArticlesCount(c *gin.Context)
	GetListByTitle(c *gin.Context)
	GetCountByTitle(c *gin.Context)
	GetListByCategory(c *gin.Context)
	GetCountByCategory(c *gin.Context)
	GetListByUser(c *gin.Context)
	GetCountByUser(c *gin.Context)
	UpdateArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)
	UserIsLiked(c *gin.Context)
	UserLikesArticle(c *gin.Context)
	UserIsStared(c *gin.Context)
	UserStarsArticle(c *gin.Context)
	GetStaredArticles(c *gin.Context)
	GetStaredArtCount(c *gin.Context)
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
		data.UserID = c.MustGet("user_id").(uint)
		code = ac.articleService.CreateArticle(&data)
	}

	var responseData *dto.ArticleResponse
	if code == errmsg.SUCCESS {
		responseData = dto.ArticleToResponse(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, code := ac.articleService.GetArticleInfo(uint(id))
	var responseData *dto.ArticleResponse
	if code == errmsg.SUCCESS {
		responseData = dto.ArticleToResponse(article)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetArticleList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//title := c.Query("title")

	var articles []model.Article
	var code int

	articles, code = ac.articleService.GetArticleList(pageSize, pageNum)
	//if len(title) == 0 {
	//	articles, code = ac.articleService.GetArticleList(pageSize, pageNum)
	//} else {
	//	articles, code = ac.articleService.GetListByTitle(title, pageSize, pageNum)
	//}
	responseData := dto.ArticleSliceToResponse(articles)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"meesage": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetAllArticlesCount(c *gin.Context) {
	total, code := ac.articleService.GetAllArticlesCount()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    total,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetListByTitle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")
	if len(title) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_BAD_REQUEST,
			"data":    nil,
			"meesage": errmsg.GetErrMsg(errmsg.ERROR_BAD_REQUEST),
		})
		return
	}
	articles, code := ac.articleService.GetListByTitle(title, pageSize, pageNum)
	responseData := dto.ArticleSliceToResponse(articles)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"meesage": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetCountByTitle(c *gin.Context) {
	title := c.Query("title")
	total, code := ac.articleService.GetArticlesCountByTitle(title)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    total,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetListByCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	articles, code := ac.articleService.GetListByCategory(uint(id), pageSize, pageNum)

	responseData := dto.ArticleSliceToResponse(articles)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetCountByCategory(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Param("id"))
	total, code := ac.articleService.GetArticlesCountByCategory(uint(categoryID))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    total,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetListByUser(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	articles, code := ac.articleService.GetListByUser(uint(id), pageSize, pageNum)

	responseData := dto.ArticleSliceToResponse(articles)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetCountByUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	total, code := ac.articleService.GetArticlesCountByUser(uint(userID))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    total,
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
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    nil,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	var requester model.User
	requester.ID = c.MustGet("user_id").(uint)
	requester.Role = c.MustGet("role").(uint8)

	code = ac.articleService.UpdateArticle(&requester, uint(id), &data)
	var responseData *dto.ArticleResponse
	if code == errmsg.SUCCESS {
		responseData = dto.ArticleToResponse(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var requester model.User

	requester.ID = c.MustGet("user_id").(uint)
	requester.Role = c.MustGet("role").(uint8)

	code := ac.articleService.DeleteArticle(&requester, uint(id))

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) UserIsLiked(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("user_id").(uint)

	data, code := ac.articleService.UserIsLiked(uint(id), userID)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) UserLikesArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("user_id").(uint)

	code := ac.articleService.UserLikesArticle(uint(id), userID)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) UserIsStared(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("user_id").(uint)

	data, code := ac.articleService.UserIsStared(uint(id), userID)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) UserStarsArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("user_id").(uint)

	code := ac.articleService.UserStarsArticle(uint(id), userID)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *ArticleController) GetStaredArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	userID := c.MustGet("user_id").(uint)

	articles, code := ac.articleService.GetStaredArticles(userID, pageSize, pageNum)

	responseData := dto.ArticleSliceToResponse(articles)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}
func (ac *ArticleController) GetStaredArtCount(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	total, code := ac.articleService.GetStaredArtCount(userID)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    total,
		"message": errmsg.GetErrMsg(code),
	})
}
