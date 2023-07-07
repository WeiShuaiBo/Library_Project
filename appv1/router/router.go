package router

import (
	"Library-management/appv1/logic"
	"Library-management/appv1/tools"
	_ "Library-management/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func New() *gin.Engine {
	r := gin.Default()
	//生成swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//所有的都可以进行查看书籍，注册登录等
	all := r.Group("/all")
	{
		all.POST("/login", logic.Login)               //登录
		all.POST("/enroll", logic.Enroll)             //注册
		all.GET("/getBook/:book_name", logic.GetBook) //模糊查书
		all.GET("/getBooks", logic.GetBooks)          //显示所有书记
	}
	//中间件，在测试swagger时可以暂时不用中间件
	all.Use(AuthCheck())
	//用户借书还书
	user := all.Group("")
	{
		user.POST("/user", logic.GetUser)
		user.PUT("/user", logic.UpdateUser)
		user.POST("/borrow/:book_name", logic.Borrow)
		user.POST("/return/:book_name", logic.Return)
		user.GET("/logout", logic.Logout)
		user.GET("/borrowing", logic.Borrowing)
	}

	//管理员
	administrator := all.Group("/ad")
	{
		//查看用户的借阅情况
		administrator.GET("/user/:name", logic.GetUserBook)
		administrator.GET("/users", logic.GetUserBooks)

		//图书的增删
		administrator.POST("/book", logic.Create)
		administrator.DELETE("/book/:book_name", logic.Delete)
		//图书的借阅信息
		administrator.GET("book/:book_name", logic.GetBookUser)
	}
	return r
}
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//测试模式不需要验签,需要自己在请求的头部伪造一个Debug数据
		if c.GetHeader("debug") != "" {
			c.Next()
			return
		}
		auth := c.GetHeader("Authorization")
		fmt.Println(auth)
		data, err := tools.Token.VerifyToken(auth[7:])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.NotLogin,
				Message: "验签失败！",
			})
		}
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
