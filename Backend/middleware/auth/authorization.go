package auth

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/utils/errmsg"
	"net/http"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var role uint8
		if roleVal, exists := c.Get("role"); exists {
			role = roleVal.(uint8)
		}
		if role == 1 {
			c.Next()
		} else {
			code := errmsg.ERROR_USER_NOT_ADMIN
			c.JSON(http.StatusForbidden, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
		}
	}
}
