package v1

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/model"
	"myblog.backend/repository"
	"myblog.backend/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var code int
	err := c.ShouldBindJSON(&data)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
	code = repository.AddUser(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// Todo 查询单个用户

// 查询所有用户
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := repository.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := repository.EditUser(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := repository.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
