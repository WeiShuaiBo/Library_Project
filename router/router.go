package router

import (
	_ "LibraryTest/docs"
	"LibraryTest/logic"
	"LibraryTest/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func New() *gin.Engine {
	router := gin.Default()
	//图书查询功能
	{
		//查看指定图书信息
		router.GET("/:book_name", logic.GetBook)
		//查看全部图书信息
		router.GET("/", logic.GetAllBook)
	}
	user := router.Group("user")
	user.Use(AuthCheck())
	{
		//用户查看个人信息
		user.GET("/", logic.GetUser)
		//用户修改个人信息
		user.PUT("/", logic.PutUser)
		//用户查看自己的借阅信息
		user.GET("/book", logic.GetMyBook)
		//用户借阅图书
		user.POST("/book/:book_id", logic.Borrow)
		////用户归还图书
		user.GET("/book/:book_id", logic.GiveBack)
	}
	//管理员功能
	manager := router.Group("/manager")

	{
		//查询所有用户信息
		manager.GET("/user", logic.GetAllUser)
		//查看指定用户的借阅信息
		manager.GET("/user/:id", logic.GetUserBorrow)
		//查看所有图书的借阅信息
		manager.GET("/", logic.GetALlBookBorrow)
		//添加图书
		manager.POST("/book", logic.AddBook)
		//删除图书
		manager.DELETE("/book/:book_id", logic.DeleteBook)
	}
	//用户登录界面
	login := router.Group("log")
	{
		//用户登录
		login.POST("/login", logic.Login)
		//用户退出
		login.GET("/logout", logic.Logout)
		//用户注册
		login.POST("/register", logic.Register)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//测试模式不需要验签,需要自己在请求的头部伪造一个Debug数据
		if c.GetHeader("debug") != "" {
			c.Next()
			return
		}
		auth := c.GetHeader("Authorization")
		auth = auth[7:]
		fmt.Printf("auth:%+v\n", auth)
		data, err := tools.Token.VerifyToken(auth)
		if err != nil {
			//终止请求并返回json响应
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.UnLoginErr,
				Message: "验签失败！",
			})
		}
		fmt.Printf("data:%+v\n", data)
		if data.ID <= 0 || data.Name == "" {
			//如果用户没有登录，中间件直接返回，不再向后继续
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.UnLoginErr,
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
