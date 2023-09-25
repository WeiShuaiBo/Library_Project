package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv11/pkg/result"
	"library/appV2/global"
	"library/appV2/logic"
	"library/appV2/model"
	"library/appV2/tools"
	"net/http"
)

func userRouter(r *gin.Engine) {
	//路由就是 /user/users
	base := r.Group("/user")
	base.Use(AuthCheck())
	user := base.Group("/users") //为什么双层分组
	{
		user.POST("/records/:bookId", logic.BorrowBook) //借书
		user.PUT("/records/:bookId", logic.ReturnBook)  //还书
		user.GET("/GetPersonalInformation/", logic.GetPersonalInformation)
		user.POST("/UpdatePersonalInformation/", logic.UpdatePersonalInformation)
		user.POST("/buy/:bookId", logic.BuyBook)
		user.POST("/qrcodeLogin/:code", logic.QrcodeLogin)
		user.GET("/aLiPay/:ordersId", logic.ALiPay)

	}
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
		if data.ID <= 0 || data.Name == "" || data.Role == "" {
			//如果用户没有登录，中间件直接返回，不再向后继续
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.HttpCode{
				Code:    tools.NotLogin,
				Message: "用户信息获取错误",
			})
			return
		}
		order := model.NoticePay(data.ID)
		if order != nil {
			c.JSON(http.StatusOK, tools.HttpCode{
				Code:    tools.UserInfoErr,
				Message: "请尽快支付，否则订单将自动取消",
				Data:    order,
			})
		}
		//rbac权限管理开始
		// 请求的path
		p := c.Request.URL.Path
		// 请求的方法
		m := c.Request.Method

		role := data.Role
		//role:="user"
		//role:="guest"

		fmt.Println("role:" + role)
		fmt.Println("path:" + p)
		fmt.Println("method:" + m)

		// 检查用户权限
		isPass, err := global.Enforcer.Enforce(role, p, m)
		if err != nil {
			resultRes := result.NewResult(c)
			fmt.Println("检测错误")
			resultRes.Error(2005, err.Error())
			return
		}
		fmt.Print(isPass)
		if isPass {
			//c.Next()
			fmt.Println("鉴权成功")
		} else {
			resultRes := result.NewResult(c)
			resultRes.Error(2006, "无访问权限")
			return
		}
		//rbac权限管理结束
		//将用户信息塞到Context中
		c.Set("name", data.Name)
		c.Set("userId", data.ID)
		c.Set("rose", data.Role)
		c.Next()
	}
}
