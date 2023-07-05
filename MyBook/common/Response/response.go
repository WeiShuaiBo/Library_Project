// @Author	zhangjiaozhu 2023/7/3 19:25:00
package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":  400,
		"error": msg,
	})
}
