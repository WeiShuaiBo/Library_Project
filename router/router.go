package router

import (
	"Library_Project/dao/mysql"
	_ "Library_Project/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
)
import "github.com/swaggo/gin-swagger"

func Router() *gin.Engine {
	r := gin.New()
	//API 接口文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//公共 可以查书 登录注册
	v1 := r.Group("/api/v1/public")
	v1.POST("/login", mysql.Login)
	v1.POST("/register", mysql.Register)
	v1.GET("/books", mysql.BGetBooks)
	v1.GET("/book/:id", mysql.BGetBook)
	v1.POST("/logout", mysql.Logout)
	//此处一个限制条件中间件，验证是否登录，未登录去登录
	v1.Use() //用来判断 是否登录限制
	//普通用户  借书还书
	user := v1.Group("")
	{
		user.GET("/selfInformation", mysql.GetUser)
		user.PUT("/updateInformation", mysql.Update)
		user.POST("/borrow/:id", mysql.Borrow)
		user.POST("/return/:id", mysql.Return)
		user.GET("/selfBorrowHistory", mysql.GetSelfBorrowHistory)
	}
	//此处一个限制中间件，验证用户是否是管理员，不是管理员不进行下面操作
	//超级管理员功能  可以增删改查用户，增删改书
	suser := v1.Group("")
	{
		//对用户信息管理  增删改查
		//suser.POST("/users", mysql.Create)
		//suser.PUT("/users/:id", mysql.Update)
		//suser.DELETE("/user/:id", mysql.Delete)
		suser.GET("/users", mysql.GetUsers)
		suser.GET("/users/:id", mysql.GetUser)
		//对图书信息管理  增删改
		suser.POST("/book", mysql.BCreate)
		//suser.PUT("/book/:id", mysql.BUpdate)
		suser.DELETE("/book/:id", mysql.BDeleteSome)
		suser.GET("/borrowHistory", mysql.BorrowHistoryAll)
		suser.GET("/borrowHistory/:id", mysql.BorrowHistorySome)
	}
	return r
}
