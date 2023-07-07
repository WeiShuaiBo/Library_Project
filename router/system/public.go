package system

import (
	"Library_Project/controller"
	"github.com/gin-gonic/gin"
)

type PublicRouter struct {
}

func (s *PublicRouter) InitPublicRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	sysRouter := Router.Group("public")

	//获取路由函数
	var loginController = controller.ApiGroupApp.SystemApiGroup
	{
		sysRouter.POST("/login", loginController.Login)
		sysRouter.POST("/register", loginController.Register)
		sysRouter.GET("/books", loginController.Books)
		sysRouter.GET("/search", loginController.Search)

	}
	return sysRouter
}
