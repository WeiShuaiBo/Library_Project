package main

import "Library_Project/router"

// @title 图书管理接口文档
// @version v1
// @description 有关图书系统相关的接口文档
// @license.name Apache 2.0
// @host 127.0.0.1:8083
// @BasePath /api/v1
func main() {
	// 基础配置

	// 启动路由
	router.Start()
}
