package v1

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/model"
	"myblog.backend/repository"
	"myblog.backend/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	err := c.ShouldBindJSON(&data)
	code := errmsg.SUCCESS
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
	} else {
		code = repository.CreateCate(&data)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询分类信息
func GetCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := repository.GetCateInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total := repository.GetCate(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 编辑分类名
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := repository.EditCate(id, &data)
	//if code == errmsg.ERROR_CATENAME_USED {
	//	c.Abort()
	//}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := repository.DeleteCate(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
