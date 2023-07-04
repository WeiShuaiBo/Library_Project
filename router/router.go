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
		router.GET("/:book_name", logic.GetBook)
		router.GET("/", logic.GetAllBook)
	}
	user := router.Group("user")
	//user.Use(AuthCheck())
	{
		//用户查看个人信息
		user.GET("/", logic.GetUser)
		//用户修改个人信息
		user.PUT("/", logic.PutUser)
		//用户查看自己的借阅信息
		user.GET("/book", logic.GetMyBook)
		//用户借阅图书
		user.POST("/book/:book_name", logic.Borrow)
		//用户归还图书
		user.GET("/book/:book_name", logic.GiveBack)
	}
	//管理员功能
	manager := router.Group("/manager")
	{
		//查询所有用户信息
		manager.GET("/user", logic.GetAllUser)
		//查看指定用户的借阅信息
		manager.GET("/user/:id")
		//查看所有图书的借阅信息
		manager.GET("/", logic.GetALlBookBorrow)
		//添加图书
		manager.POST("/book", logic.AddBook)
		//删除图书
		manager.DELETE("/book/:book_name", logic.DeleteBook)
	}
	//用户登录界面
	login := router.Group("")
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
		//用于测试，免验签
		if c.GetHeader("Debug") != "" {
			c.Next()
			return
		}
		auth := c.GetHeader("authorization")
		fmt.Println("authorization:", auth)
		data, err := tools.Token.VerifyToken(auth)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.UnLoginErr,
				Message: "验签失败",
			})
		}
		fmt.Println("data:", data)
		if data.ID <= 0 || data.Name == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.UnFoundErr,
				Message: "用户信息获取失败",
			})
			return
		}
		c.Set("name", data.Name)
		c.Set("userId", data.ID)
		c.Next()

	}
}
