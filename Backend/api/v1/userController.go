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

type IUserController interface {
	CreateUser(c *gin.Context)
	GetUserList(c *gin.Context)
	UpdateUserBasicInfo(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	userService := service.NewUserService()
	return &UserController{userService}
}

/* ====================================== */

// 新增用户
func (uc *UserController) CreateUser(c *gin.Context) {
	var user model.User
	var code int
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
	code = uc.userService.CreateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表（分页）
func (uc *UserController) GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	users, total, code := uc.userService.GetUserList(pageSize, pageNum)
	responseData := dto.UserSliceToResponse(users)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户基本信息（只限于FullName, Bio）
func (uc *UserController) UpdateUserBasicInfo(c *gin.Context) {
	var user model.User
	var code int
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}

	code = uc.userService.UpdateUserBasicInfo(uint(id), &user)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func (uc *UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := uc.userService.DeleteUser(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
