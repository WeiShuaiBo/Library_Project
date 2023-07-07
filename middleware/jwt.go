package middleware

import (
	"Library_Project/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWTAuth
// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息
// 这里前端需要把token存储到cookie或者本地localStorage1
// 可以约定刷新令牌或者重新登录
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})

			//abort()函数执行终止当前中间件以后的中间件执行，但是会执行当前中间件的后续逻辑
			c.Abort()
			return

		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := pkg.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的UserId信息保存到请求的上下文c上
		c.Set("UserId", mc.UserId)

		//next()函数会跳过当前中间件中next()后的逻辑，当下一个中间件执行完成后再执行剩余的逻辑
		c.Next() // 后续的处理函数可以用过c.Get("UserId")来获取当前请求的用户信息
	}

}
