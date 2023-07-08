package appV0

import (
	"library/appV0/logger"
	"library/appV0/model"
	"library/appV0/router"
	"library/appV0/tools"
)

// @title			图书管理系统
// @version		0.0
// @description	图书管理系统，现在是0.0版本
func Start() {
	logger.InitLogger()
	defer func() {
		model.Close()
	}()
	r := router.New()

	model.New()
	tools.NewToken("", "")

	_ = r.Run(":8080")

}
