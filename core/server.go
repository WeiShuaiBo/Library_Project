package core

import (
	"Library_Project/global"
	"Library_Project/initialize"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.FAST_CONFIG.System.Port)

	// 初始化swagger
	initialize.Swagger(Router)
	s := initServer(address, Router)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.FAST_LOG.Info("server run success on", zap.String("address", address))
	fmt.Printf(`	欢迎使用 FAST-GIN
		当前版本: V1.0
		默认自动化文档地址:http://127.0.0.1%s/api/swagger/index.html
	`, address)
	global.FAST_LOG.Error(s.ListenAndServe().Error())
}
