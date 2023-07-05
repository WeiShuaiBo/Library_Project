package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "library/docs"
	"library/logger"
	"library/logic"
	"library/tools"
	"net/http"
)

func New() *gin.Engine {

	r := gin.Default()
	{
		r.GET("/books", logic.GetBooks)                  //查询所有图书
		r.POST("/bookByKeyWord", logic.GetBookByKeyWord) //根据图书名模糊查询
		r.POST("/login", logic.Login)                    //用户登录
		r.POST("/register", logic.Register)              //用户注册
		r.POST("/adminLogin", logic.AdminLogin)          //管理员登录
	}

	admin := r.Group("/admin")
	admin.Use(AdminJWTMiddleware())
	{
		admin.GET("/books", logic.AdminGetBooks)                  //管理员查询所有图书详细信息
		admin.POST("/bookByKeyWord", logic.AdminGetBookByKeyWord) //管理员根据图书名模糊查询详细信息
		admin.GET("/book/:id", logic.AdminGetBooksById)           //管理员根据id查询图书详细信息
		admin.POST("/book", logic.CreatBook)                      //管理员新建图书
		admin.PUT("/book/:id", logic.UpdateBook)                  //管理员修改图书
		admin.DELETE("/book/:id", logic.DeleteBook)               //管理员删除图书
		admin.GET("/getInfo", logic.AdminGetInfo)                 //管理员查询所有借阅表详细信息
	}

	user := r.Group("/user")
	user.Use(AuthCheck())
	{
		user.GET("/getInfo", logic.GetInfo)              //用户获取个人信息
		user.PUT("/updateInfo", logic.UpdateInfo)        //用户修改个人信息
		user.PUT("/lendBook/:id", logic.UserLendBook)    //用户借书
		user.PUT("/giveBook/:id", logic.UserGiveBook)    //用户还书
		user.GET("getAllLendInfo", logic.GetAllLendInfo) //用户获取个人借阅记录
	}

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

// AuthCheck 用户验证中间件
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		logger.Log.Info("auth：", auth)
		data, err := tools.Token.VerifyToken(auth)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.NotLogin,
				Message: "验签失败",
			})
			return
		}
		logger.Log.Info("data", data)
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
	}
}

// AdminJWTMiddleware 管理员验证中间件
func AdminJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		logger.Log.Info("auth：", auth)
		data, err := tools.Token.AdminVerifyToken(auth)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.AdminNotLogin,
				Message: "管理员验签失败",
			})
			return
		}
		logger.Log.Info("data", data)
		if data.ID <= 0 || data.Name == "" {
			//如果管理员没有登录，中间件直接返回，不再向后继续
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.NotLogin,
				Message: "管理员信息获取错误",
			})
			return
		}
		//将用户信息塞到Context中
		c.Set("name", data.Name)
		c.Set("adminId", data.ID)
	}
}
