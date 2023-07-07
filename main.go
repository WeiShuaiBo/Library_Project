package main

import (
	"Library_Project/core"
	"Library_Project/global"
	"Library_Project/initialize"
	"fmt"
	"go.uber.org/zap"
	"os/exec"
)

// @title                       图书管理
// @version                     0.0.1
// @description                测试
// host							127.0.0.1:8181
func main() {
	cmd := exec.Command("swag", "init")
	fmt.Println("重新加载swagger成功", cmd.Args)
	global.FAST_VP = initialize.Viper()
	global.FAST_LOG = initialize.Zap()
	zap.ReplaceGlobals(global.FAST_LOG)
	global.FAST_DB = initialize.Gorm()
	//if global.FAST_DB != nil {
	//	//todo: 初始化数据库
	//	initialize.InitTables(global.FAST_DB)
	//initialize.InitDBSource()
	//
	//	db, _ := global.FAST_DB.DB()
	//	defer func(db *sql.DB) {
	//		db.Close()
	//	}(db)
	//}

	core.RunWindowsServer()
}
