// @Author	zhangjiaozhu 2023/7/4 20:18:00
package midwares

import (
	"MyBook/common/Response"
	"MyBook/models"
	"MyBook/models/DB/MysqlDB"
	"github.com/gin-gonic/gin"
)

func Admin() func(c *gin.Context) {
	return func(c *gin.Context) {
		value, exists := c.Get(ContextUserIDKey)
		if !exists {
			c.JSON(200, "未登录")
		}
		u := value.(uint64)
		var user models.User
		MysqlDB.DB.Where("user_id= ?", u).First(&user)
		if user.Role == 0 {
			Response.Error(c, "非管理员，非法操作")
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
