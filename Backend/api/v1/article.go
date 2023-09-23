package v1

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/model"
	"myblog.backend/repository"
	"myblog.backend/utils/errmsg"
	"net/http"
	"strconv"
)

// 新增文章
func AddArticle(c *gin.Context) {
	var data model.Article
	var code int
	err := c.ShouldBindJSON(&data)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		code = repository.CreateArt(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	if pageSize >= 100 {
		pageSize = 100
	} else if pageSize <= 0 {
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := repository.GetCateArt(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询文章
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := repository.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询文章列表
func GetArtList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	if len(title) == 0 {
		data, code, total := repository.GetArt(pageSize, pageNum)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		data, code, total := repository.SearchArt(title, pageSize, pageNum)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"meesage": errmsg.GetErrMsg(code),
		})
	}
}

// 编辑文章
func EditArt(c *gin.Context) {
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

	code = repository.EditArt(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := repository.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
