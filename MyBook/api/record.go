// @Author	zhangjiaozhu 2023/7/5 9:00:00
package api

import (
	"MyBook/common/Response"
	"MyBook/routers/midwares"
	"MyBook/service"
	"github.com/gin-gonic/gin"
)

func FindUserRecord(c *gin.Context) {
	// 获取当前用户的id
	value, exists := c.Get(midwares.ContextUserIDKey)
	if !exists {
		Response.Error(c, "未登录")
		return
	}
	u := value.(uint64)
	record, err := service.FindUserRecord(u)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "success", record)
}
func GetAllRecord(c *gin.Context) {
	record, err := service.GetAllRecord()
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "success", record)
}

func GetRecordByUser(c *gin.Context) {
	type Req struct {
		Username string `json:"username"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	user, err := service.GetRecordByUser(req.Username)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "success", user)
}
