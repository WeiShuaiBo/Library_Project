package initialize

import (
	"Library_Project/global"
	"Library_Project/initialize/internal"
	"Library_Project/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Zap() (logger *zap.Logger) {

	if ok, _ := utils.PathExists(global.FAST_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.FAST_CONFIG.Zap.Director)
		// Mkdir() 创建一个文件夹 指定名称和操作限权
		_ = os.Mkdir(global.FAST_CONFIG.Zap.Director, os.ModePerm)
	}

	// cores 设置日志库格式
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.FAST_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}
