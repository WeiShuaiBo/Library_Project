package system

import (
	"Library_Project/controller"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	sysRouter := Router.Group("sys")
	//获取路由函数
	var loginController = controller.ApiGroupApp.SystemApiGroup
	{
		sysRouter.GET("/information", loginController.Information)
		sysRouter.POST("/changeinfo", loginController.ChangeInfo)
		sysRouter.POST("/borrow", loginController.Borrow)
		sysRouter.POST("/returnbook", loginController.ReturnBook)
		sysRouter.GET("/records", loginController.Records)
		sysRouter.GET("/exit", loginController.Exit)

	}
	return sysRouter
}
