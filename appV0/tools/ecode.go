package tools

const (
	OK            = 0
	NotLogin      = 10001
	AdminNotLogin = 10002
	UserInfoErr   = 10003
	BookErr       = 10004
	AdminInfoErr  = 10005
	NotFound      = 10006
)

type HttpCode struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
