package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"library/appV0/config"
	_ "library/appV0/docs"
	"library/appV0/logger"
	"library/appV0/logic"
	"library/appV0/tools"
	"net/http"
)

func New() *gin.Engine {

	gin.SetMode(config.Conf.Mode)

	r := gin.Default()
	// 注册zap相关中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	guest := r.Group("guest")
	{
		guest.POST("/login", logic.Login)
		guest.POST("/register", logic.Register)
		guest.POST("/getValidateCode", logic.GetValidateCode)
		guest.POST("/findLibrary", logic.FindLibrary)
		guest.POST("/getLibrary", logic.GetLibrary)
	}

	user := r.Group("user")
	user.Use(AuthCheck())
	{

		user.GET("/logout", logic.Logout)

		common := user.Group("common")
		{
			common.POST("/dueLibrary", logic.DueLibrary)
			common.GET("/getRecordByUserId", logic.GetRecordByUserId)
			common.POST("/loanLibrary", logic.LoanLibrary)
			common.PUT("/putUserPwd", logic.PutUserPwd)
		}

		admin := user.Group("admin")
		{
			admin.POST("/creatLibrary", logic.CreatLibrary)
			admin.DELETE("/deleteLibrary", logic.DeleteLibrary)
			admin.POST("/getARecordByUserId", logic.GetARecordByUserId)
			admin.GET("/getExpected", logic.GetExpected)
			admin.GET("/getRecord", logic.GetRecord)
			admin.PUT("/updateLibrary", logic.UpdateLibrary)
		}
	}

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		fmt.Printf("auth:%+v\n", auth)
		data, err := tools.Token.VerifyToken(auth)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.NotLogin,
				Message: "验签失败！",
			})
		}
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
