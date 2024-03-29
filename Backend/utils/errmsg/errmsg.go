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

	// code=2000... 文章模块错误
	ERROR_ARTICLE_NOT_EXIST   = 2001
	ERROR_ARTICLE_TITLE_EXIST = 2002

	// code=3000... 分类模块错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002

	// code=4000... 评论模块错误
	ERROR_COMMENT_NOT_EXIST = 4001

	// code=5000... 注册模块错误
	ERROR_BAD_USERNAME        = 5001
	ERROR_BAD_PASSWORD        = 5002
	ERROR_PASSWORDS_NOT_EQUAL = 5003

	// code=6000... 上传文件模块错误
	ERROR_UPLOAD_USERAVT = 6001

	// code=7000... Redis模块错误
	REDIS_ERROR            = 7000
	REDIS_SET_NOT_EXISTS   = 7001
	REDIS_SET_ISNOT_MEMBER = 7002
	REDIS_SET_IS_MEMBER    = 7003
	REDIS_IS_SYNCING       = 7004
	REDIS_LIST_NOT_EXISTS  = 7004
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "Fail",

	// 请求错误
	ERROR_BAD_REQUEST: "请求格式错误",
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

	// 文章模块错误
	ERROR_ARTICLE_NOT_EXIST:   "文章不存在",
	ERROR_ARTICLE_TITLE_EXIST: "文章标题已存在",

	// 分类模块错误
	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",

	// 评论模块错误
	ERROR_COMMENT_NOT_EXIST: "该评论不存在",

	// 注册模块错误
	ERROR_PASSWORDS_NOT_EQUAL: "两次密码输入不一致",
	ERROR_BAD_USERNAME:        "用户名只能是字母和数字，长度为8-25",
	ERROR_BAD_PASSWORD:        "密码只能是字母、数字和特殊字符，长度为8-25",

	// 上传文件模块错误
	ERROR_UPLOAD_USERAVT: "头像上传失败",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
