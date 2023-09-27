package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"myblog.backend/config"
	"myblog.backend/utils/errmsg"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte(config.JwtKey)

type MyCustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     uint8  `json:"role"`
	jwt.RegisteredClaims
}

// 生成Token
func GenerateToken(userID uint, username string, role uint8) (string, int) {
	expireTime := time.Now().Add(24 * time.Hour)
	setClaims := MyCustomClaims{
		userID,
		username,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "SixGod",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	signedString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return signedString, errmsg.SUCCESS
}

// 验证Token
func CheckToken(signedString string) (*MyCustomClaims, int) {
	token, err := jwt.ParseWithClaims(signedString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, errmsg.ERROR
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, errmsg.SUCCESS
	}
	return nil, errmsg.ERROR
}

// JWT中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		var code int
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		claims, code := CheckToken(checkToken[1])
		if code == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > (*claims.ExpiresAt).Time.Unix() {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}
