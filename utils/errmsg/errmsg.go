package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	//code=1000 ...用户模块错误

	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenExist     = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight    = 1008
	//code=2000 ...文章模块错误

	ErrorArtNotExits = 3001

	//code=3000 ...分类模块错误

	ErrorCateNameUsed = 2001
	ErrorCateNotExits = 2002
)

var codeMsg = map[int]string{
	SUCCESS:             "ok",
	ERROR:               "FAIL",
	ErrorUsernameUsed:   "用户名已存在",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户名不存在",
	ErrorTokenExist:     "TOKEN不存在",
	ErrorTokenRuntime:   "TOKEN已过期",
	ErrorTokenWrong:     "TOKEN不正确",
	ErrorTokenTypeWrong: "TOKEN格式不正确",
	ErrorUserNoRight:    "该用户无权限",
	ErrorCateNameUsed:   "该分类已存在",
	ErrorCateNotExits:   "该分类不存在",

	ErrorArtNotExits: "文章不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
