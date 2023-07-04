package main

import (
	"LibraryTest/dao"
	"LibraryTest/router"
)

// @title 		图书馆管理系统
// @version 	1.0
// @description 该项目用户设有游客，普通用户和管理员，但是只有普通用户和管理员可以借阅书籍

// @host 		127.0.0.1
// @BasePath 	/
func main() {
	dao.New()
	defer func() {
		dao.Close()
	}()
	r := router.New()
	_ = r.Run(":80")
}
