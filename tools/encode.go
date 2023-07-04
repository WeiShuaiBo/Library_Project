package tools

const (
	OK         = 0     //成功
	DOErr      = 10001 //操作错误
	UnFoundErr = 10002 //未找到
	UnLoginErr = 10003 //用户未登录
)

type HttpCode struct {
	Code    int
	Message string
	Data    interface{}
}
