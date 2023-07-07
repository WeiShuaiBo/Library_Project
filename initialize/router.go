package initialize

import (
	//_ "Library_Project/docs"
	"Library_Project/global"
	"Library_Project/middleware"
	"Library_Project/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
// todo 初始化总路由
func Routers() *gin.Engine {
	Router := gin.New()

	Router.Use(middleware.GinLogger())

	//Router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//加载文件
	// 全局注册路由组实例位置
	// todo:  注册路有组实例
	systemRouter := router.RouterGroupApp.System
	//    公开路有组,不做权限鉴定
	PublicGroup := Router.Group("api")
	PublicGroup.Use()
	{
		systemRouter.InitPublicRouter(PublicGroup)

	}
	//    私有路有组有拦截
	PrivateGroup := Router.Group("api")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitAdminRouter(PrivateGroup)
	}
	global.FAST_LOG.Info("router register success")

	return Router
}
