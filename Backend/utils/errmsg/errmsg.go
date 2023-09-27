package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_BAD_REQUEST = 400

	// code=1000... 用户模块错误
	ERROR_USERNAME_USED       = 1001
	ERROR_PASSWORD_WRONG      = 1002
	ERROR_USER_NOT_EXIST      = 1003
	ERROR_TOKEN_NOT_EXIST     = 1004
	ERROR_TOKEN_RUNTIME       = 1005
	ERROR_TOKEN_WRONG         = 1006
	ERROR_TOKEN_TYPE_WRONG    = 1007
	ERROR_USER_NO_RIGHT       = 1008
	ERROR_USER_FULLNAME_EXIST = 1009
	ERROR_USER_NOT_ADMIN      = 1010
	ERROR_PASSWORDS_NOT_EQUAL = 1011

	// code=2000... 文章模块错误
	ERROR_ARTICLE_NOT_EXIST   = 2001
	ERROR_ARTICLE_TITLE_EXIST = 2002

	// code=3000... 分类模块错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002

	// code=4000... 评论模块错误
	ERROR_COMMENT_NOT_EXIST = 4001
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "Fail",

	// 请求错误
	ERROR_BAD_REQUEST: "请求错误",
	// 用户模块错误
	ERROR_USERNAME_USED:       "账号已存在",
	ERROR_PASSWORD_WRONG:      "密码错误",
	ERROR_USER_NOT_EXIST:      "用户不存在",
	ERROR_TOKEN_NOT_EXIST:     "TOKEN不存在，请重新登录",
	ERROR_TOKEN_RUNTIME:       "TOKEN已过期，请重新登录",
	ERROR_TOKEN_WRONG:         "TOKEN不正确，请重新登录",
	ERROR_TOKEN_TYPE_WRONG:    "TOKEN格式错误，请重新登录",
	ERROR_USER_NO_RIGHT:       "该用户无权限",
	ERROR_USER_FULLNAME_EXIST: "用户昵称已存在",
	ERROR_USER_NOT_ADMIN:      "非管理员用户无权操作",
	ERROR_PASSWORDS_NOT_EQUAL: "两次密码输入不一致",

	// 文章模块错误
	ERROR_ARTICLE_NOT_EXIST:   "文章不存在",
	ERROR_ARTICLE_TITLE_EXIST: "文章标题已存在",

	// 分类模块错误
	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",

	// 评论模块错误
	ERROR_COMMENT_NOT_EXIST: "该评论不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
