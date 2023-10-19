package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"myblog.backend/dto"
	"myblog.backend/middleware/auth"
	"myblog.backend/model"
	"myblog.backend/service"
	"myblog.backend/utils/errmsg"
	"myblog.backend/utils/valid"
	"net/http"
	"strconv"
)

/* ====================================== */

type IUserController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetUserInfo(c *gin.Context)
	UpdateUserBasicInfo(c *gin.Context)
	UpLoadAvatar(c *gin.Context)
	GetSelfProfile(c *gin.Context)
}

type UserController struct {
	userService  *service.UserService
	minioService *service.MinIOService
}

func NewUserController() *UserController {
	userService := service.NewUserService()
	minioService := service.NewMinIOService()
	return &UserController{userService, minioService}
}

/* ====================================== */

// 登录
func (uc *UserController) Login(c *gin.Context) {
	var user model.User
	var code int
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	code = uc.userService.CheckPassword(&user)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
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

// 注册
func (uc *UserController) Register(c *gin.Context) {
	var rq dto.RegisterRequest
	var code int
	var msg string
	code = errmsg.SUCCESS
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		code = errmsg.ERROR_BAD_REQUEST
		msg = errmsg.GetErrMsg(code)
	} else if code = valid.ValidateCredentials(rq.Username, rq.Password); code != errmsg.SUCCESS {
		msg = errmsg.GetErrMsg(code)
	} else if rq.Password != rq.ConfirmPassword {
		code = errmsg.ERROR_PASSWORDS_NOT_EQUAL
		msg = errmsg.GetErrMsg(code)
	}
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    nil,
			"message": msg,
		})
		return
	}

	user := dto.RegisterRequestToUser(&rq)
	user.Role = 2
	code = uc.userService.CreateUser(user)
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

func (uc *UserController) GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, code := uc.userService.GetUserByID(uint(id))
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

// 编辑用户基本信息（只限于FullName, Bio）
func (uc *UserController) UpdateUserBasicInfo(c *gin.Context) {
	var user model.User
	var code int
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&user)
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

	code = uc.userService.UpdateUserBasicInfo(&requester, uint(id), &user)
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

// 上传用户头像
func (uc *UserController) UpLoadAvatar(c *gin.Context) {
	file, _, err := c.Request.FormFile("avatar")
	var code int
	var url string
	if err != nil {
		log.Println("文件请求错误", err)
	} else {
		userID := c.MustGet("user_id").(uint)
		url, code = uc.minioService.UpLoadUserAvatar(userID, &file)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     code,
		"message":    errmsg.GetErrMsg(code),
		"avatar_url": url,
	})
}

// 获取自己的个人信息
func (uc *UserController) GetSelfProfile(c *gin.Context) {
	selfID := c.MustGet("user_id").(uint)
	user, code := uc.userService.GetUserByID(selfID)
	responseData := dto.UserToResponse(user)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    responseData,
	})
}
