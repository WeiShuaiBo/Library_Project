package system

import (
	"Library_Project/controller"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (s *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	sysRouter := Router.Group("admin")
	//获取路由函数
	var loginController = controller.ApiGroupApp.SystemApiGroup
	{
		sysRouter.POST("/addbook", loginController.AddBook)
		sysRouter.POST("/delbook", loginController.DelBook)
		sysRouter.POST("/searchperbook", loginController.SearchPerBook)
		sysRouter.POST("/searchbook", loginController.SearchBook)
		//sysRouter.GET("/asd", loginController.SearchBook)
	}
	return sysRouter
}
