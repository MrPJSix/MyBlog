package v1

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/dto"
	"myblog.backend/middleware/auth"
	"myblog.backend/model"
	"myblog.backend/service"
	"myblog.backend/utils/errmsg"
	"net/http"
	"strconv"
)

/* ====================================== */

type IAdminController interface {
	AdminLogin(c *gin.Context)
	CreateAdmin(c *gin.Context)
	GetUserInfo(c *gin.Context)
	GetUserList(c *gin.Context)
	UpdateUserBasicInfo(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsersCount(c *gin.Context)
}

type AdminController struct {
	userService *service.UserService
}

func NewAdminController() *AdminController {
	userService := service.NewUserService()
	return &AdminController{userService}
}

/* ====================================== */

// 管理员登录
func (ac *AdminController) AdminLogin(c *gin.Context) {
	var user model.User
	var code int
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	code = ac.userService.CheckPassword(&user)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	} else if code == errmsg.SUCCESS && user.Role == 2 {
		code = errmsg.ERROR_USER_NO_RIGHT
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		var token string
		token, code = auth.GenerateToken(user.ID, user.Username, user.Role)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"token":   token,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 创建管理员账号
func (ac *AdminController) CreateAdmin(c *gin.Context) {
	var user model.User
	var code int
	err := c.ShouldBindJSON(&user)
	user.Role = 1 // 授权管理员
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"data":    nil,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	code = ac.userService.CreateUser(&user)
	var responseData *dto.UserResponse
	if code == errmsg.SUCCESS {
		responseData = dto.UserToResponse(&user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户信息
func (ac *AdminController) GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, code := ac.userService.GetUserByID(uint(id))
	var responseData *dto.UserResponse
	if code == errmsg.SUCCESS {
		responseData = dto.UserToResponse(user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表（分页）
func (ac *AdminController) GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	users, total, code := ac.userService.GetUserList(pageSize, pageNum)
	responseData := dto.UserSliceToResponse(users)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户基本信息（只限于FullName, Bio）
func (ac *AdminController) UpdateUserBasicInfo(c *gin.Context) {
	var user model.User
	var code int
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"data":    nil,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	var requester model.User
	requester.ID = c.MustGet("user_id").(uint)
	requester.Role = c.MustGet("role").(uint8)

	code = ac.userService.UpdateUserBasicInfo(&requester, uint(id), &user)
	var responseData *dto.UserResponse
	if code == errmsg.SUCCESS {
		responseData = dto.UserToResponse(&user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func (ac *AdminController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := ac.userService.DeleteUser(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func (ac *AdminController) GetAllUsersCount(c *gin.Context) {
	total, code := ac.userService.GetAllUsersCount()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    total,
		"message": errmsg.GetErrMsg(code),
	})
}
