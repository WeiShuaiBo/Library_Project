package main

import (
	"library/logger"
	"library/model"
	"library/router"
	"library/tools"
)

// @title			图书管理系统
// @version		0.0
// @description	图书管理系统，现在是0.0版本
func main() {
	logger.InitLogger()
	defer func() {
		model.Close()
	}()
	r := router.New()

	model.New()
	tools.NewToken("", "")

	_ = r.Run(":8080")

}
