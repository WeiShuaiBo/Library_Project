package router

import (
	_ "AgroWeather/app/docs"
	"AgroWeather/app/logic"
	"AgroWeather/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

func New() *gin.Engine {
	r := gin.Default()

	//weather.Use(AuthCheck()) //后续添加 登录和限流中间件

	weather := r.Group("/weather")
	{
		weather.POST("/do", logic.GetW)
	}

	login := r.Group("")
	{
		login.POST("/login", logic.Login)
		login.POST("/register", logic.Register)
		login.POST("/getValidateCode", logic.GetValidateCode)
		login.PUT("/putUserPwd", logic.PutUserPwd)

		login.GET("/logout", logic.Logout)
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
