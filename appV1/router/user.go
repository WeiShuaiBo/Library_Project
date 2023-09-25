package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/appV1/logic"
	"library/appV1/tools"
	"net/http"
)

func userRouter(r *gin.Engine) {
	//路由就是 /user/users
	base := r.Group("/user")
	base.Use(AuthCheck())
	user := base.Group("/users") //为什么双层分组
	{
		//user.GET("", logic.GetUser) //
		//user.PUT("/:id", logic.UpdateUser)
		//user.PUT("", logic.UpdateUser) //
		////user.DELETE(":id", logic.DeleteUser)
		//user.GET("/:id/records", logic.GetUserRecords)               //
		//user.GET("/:id/records/:status", logic.GetUserStatusRecords) //
		////用户自助借书还书
		user.POST("/records/:bookId", logic.BorrowBook) //借书
		user.PUT("/records/:bookId", logic.ReturnBook)  //还书
		user.GET("/GetPersonalInformation/", logic.GetPersonalInformation)
		user.POST("/UpdatePersonalInformation/", logic.UpdatePersonalInformation)

	}

	// 暂不使用
	/*book := base.Group("/books")
	{
		book.GET("/:id", logic.GetBook)
		//book.POST("/:id", logic.AddBook)
		//book.DELETE("/:id", logic.DeleteBook)
	}
	category := base.Group("/categories")
	{
		//category.GET("/:id", logic.GetCategory)
		category.GET("/:id/books/:type", logic.GetCategoryBooks)
	}*/
}

func AuthCheck() gin.HandlerFunc {
	//gin.HandlerFunc 用来处理 HTTP 请求的函数
	return func(c *gin.Context) {
		//测试模式不需要验签,需要自己在请求的头部伪造一个Debug数据
		if c.GetHeader("debug") != "" {
			c.Next()
			return
		}
		//debug是什么？
		auth := c.GetHeader("Authorization")
		fmt.Printf("auth:%+v\n", auth)
		data, err := tools.Token.VerifyToken(auth)
		//data 表示解析后的 JWT Token 信息，通常包含了一些用于身份认证、鉴权等的关键信息（如用户 ID、角色、权限等）
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.NotLogin,
				Message: "验签失败！",
			})
		}
		//c.AbortWithStatusJSON() 函数中止当前的请求并向客户端返回一个相应的错误信息。
		fmt.Printf("data:%+v\n", data)
		if data.ID <= 0 || data.Name == "" {
			//如果用户没有登录，中间件直接返回，不再向后继续
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.NotLogin,
				Message: "用户信息获取错误",
			})
			return
		}

		//将用户信息塞到Context中
		c.Set("name", data.Name)
		c.Set("userId", data.ID)

		c.Next()
	}
}
