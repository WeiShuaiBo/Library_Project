package initialize

import (
	"Library_Project/docs"
	"Library_Project/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Swagger(router *gin.Engine) {
	if global.FAST_CONFIG.Swagger.Enabled {
		address := global.FAST_CONFIG.Swagger.Host
		docs.SwaggerInfo.Title = global.FAST_CONFIG.Swagger.Title
		docs.SwaggerInfo.Description = global.FAST_CONFIG.Swagger.Description
		docs.SwaggerInfo.Version = global.FAST_CONFIG.Swagger.Version
		docs.SwaggerInfo.Host = address
		docs.SwaggerInfo.BasePath = global.FAST_CONFIG.Swagger.BasePath
		docs.SwaggerInfo.Schemes = global.FAST_CONFIG.Swagger.Schemes
		router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		global.FAST_LOG.Info("register swagger handler")
	}
}
