package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	dto2 "myblog.backend/dto/request"
	"myblog.backend/dto/response"
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
	UpdateSelfBasicInfo(c *gin.Context)
	UpLoadAvatar(c *gin.Context)
	GetSelfProfile(c *gin.Context)
	UserIsFollowed(c *gin.Context)
	UserFollow(c *gin.Context)
	GetSelfTheFollowed(c *gin.Context)
	GetSelfFans(c *gin.Context)
	GetOtherTheFollowed(c *gin.Context)
	GetOtherFans(c *gin.Context)
	GetTop5Authors(c *gin.Context)
	GetTop5FollowedUsers(c *gin.Context)
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
	var rq dto2.RegisterRequest
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

	user := dto2.RegisterRequestToUser(&rq)
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
func (uc *UserController) UpdateSelfBasicInfo(c *gin.Context) {
	var user model.User
	var code int
	id := c.MustGet("user_id").(uint)
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

	code = uc.userService.UpdateSelfBasicInfo(uint(id), &user)
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
		code = errmsg.ERROR_UPLOAD_USERAVT
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

func (uc *UserController) UserIsFollowed(c *gin.Context) {
	followedID, _ := strconv.Atoi(c.Param("id"))
	followerID := c.MustGet("user_id").(uint)

	data, code := uc.userService.UserIsFollowed(followerID, uint(followedID))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
func (uc *UserController) UserFollow(c *gin.Context) {
	followedID, _ := strconv.Atoi(c.Param("id"))
	followerID := c.MustGet("user_id").(uint)

	code := uc.userService.UserFollow(followerID, uint(followedID))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
func (uc *UserController) GetSelfTheFollowed(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	selfID := c.MustGet("user_id").(uint)

	users, code := uc.userService.GetTheFollowed(selfID, pageSize, pageNum)
	responseData := dto.UserSliceToResponse(users)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

func (uc *UserController) GetSelfFans(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	selfID := c.MustGet("user_id").(uint)

	users, code := uc.userService.GetFans(selfID, pageSize, pageNum)
	responseData := dto.UserSliceToResponse(users)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

func (uc *UserController) GetOtherTheFollowed(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	userID, _ := strconv.Atoi(c.Param("id"))

	users, code := uc.userService.GetTheFollowed(uint(userID), pageSize, pageNum)
	responseData := dto.UserSliceToResponse(users)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

func (uc *UserController) GetOtherFans(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	userID, _ := strconv.Atoi(c.Param("id"))

	users, code := uc.userService.GetFans(uint(userID), pageSize, pageNum)
	responseData := dto.UserSliceToResponse(users)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    responseData,
		"message": errmsg.GetErrMsg(code),
	})
}

func (uc *UserController) GetTop5Authors(c *gin.Context) {
	users, code := uc.userService.GetTop5Authors()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"message": errmsg.GetErrMsg(code),
	})
}
func (uc *UserController) GetTop5FollowedUsers(c *gin.Context) {
	users, code := uc.userService.GetTop5FollowedUsers()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"message": errmsg.GetErrMsg(code),
	})
}
