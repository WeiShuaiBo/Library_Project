// @Author	zhangjiaozhu 2023/7/3 14:50:00
package routers

import (
	"MyBook/api"
	"MyBook/common/config"
	_ "MyBook/docs" // 千万不要忘了导入把你上一步生成的docs
	"MyBook/routers/midwares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1/")
	{
		// 登陆/注册
		v1.POST("user/login", api.UserLogin)
		v1.POST("user/register", api.UserRegister)
		// 刷新token
		v1.GET("refreshToken", api.RefreshTokenHandler)
		// 发送邮箱
		v1.POST("user/sendEmail", api.SendEmail)
		// 游客查询书籍
		v1.GET("book/getBook", api.GetBook)                // 查询所有书籍
		v1.POST("book/getBookList")                        //分页查询书籍
		v1.POST("book/findBookByName", api.FindBookByName) //根据书名查询书籍
	}
	v1.Use(midwares.JWTAuth())
	{
		// 普通用户
		v1.POST("book/borrowBook", api.BorrowBook)         //借阅
		v1.POST("book/returnBook", api.ReturnBook)         //归还
		v1.GET("user/findUserInfo", api.FindUserInfo)      // 查询个人信息
		v1.POST("user/updateUserInfo", api.UpdateUserInfo) // 修改个人信息
		v1.GET("user/findUserRecord", api.FindUserRecord)  //查询自己的借阅记录
		v1.GET("user/logout", api.Logout)                  // 退出登录
		v1.Use(midwares.Admin())
		{
			// 管理员
			v1.POST("user/deleteUser", api.DeleteUser) // 删除用户
			v1.GET("user/getAllUser", api.GetAllUser)  // 获取所有用户
			//v1.POST("user/updateUser",api.UpdateUser) // 修改用户
			//v1.POST("user/findUserByPhone",api.FindUserByEmail) // 根据邮箱查找用户
			v1.POST("book/createBook", api.CreateBook)             //添加书籍
			v1.POST("book/deleteBookByName", api.DeleteBookByName) // 根据书名删除书籍
			v1.POST("book/updateBook", api.UpdateBook)             // 更新书籍
			// 查询所有借阅书籍记录
			v1.GET("record/getAllRecord", api.GetAllRecord)
			// 根据用户查询书籍记录
			v1.POST("record/getRecordByUser", api.GetRecordByUser)
		}
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	err := r.Run(config.Config.System.HttpPort)
	if err != nil {
		panic(err)
	}
	return r
}
