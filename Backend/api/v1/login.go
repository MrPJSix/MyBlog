package v1

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/middleware/auth"
	"myblog.backend/model"
	"myblog.backend/service"
	"myblog.backend/utils/errmsg"
	"net/http"
)

/* ====================================== */

type ILoginController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	CreateAdminUser(c *gin.Context)
}

type LoginController struct {
	userService *service.UserService
}

func NewLoginController() *LoginController {
	userService := service.NewUserService()
	return &LoginController{userService}
}

/* ====================================== */

func (lc *LoginController) Login(c *gin.Context) {
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
	code = lc.userService.CheckPassword(&user)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		var token string
		token, code = auth.GenerateToken(user.Username, user.Role)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"token":   token,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

func (lc *LoginController) Register(c *gin.Context) {

}

func (lc *LoginController) CreateAdminUser(c *gin.Context) {

}
