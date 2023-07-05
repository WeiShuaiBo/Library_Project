package tools

const (
	OK            = 0
	NotLogin      = 10001 //您还没有登录
	UserInfoErr   = 10002 //用户信息错误
	UserRepeatErr = 10004 //用户名重复
	CaptchaNil    = 10005 //验证码为空
	DoErr         = 10003

	NotFound = 10004 //信息不存在
)

type HttpCode struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
