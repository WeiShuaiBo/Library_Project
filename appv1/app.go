package appv1

import (
	"Library-management/appv1/modle"
	"Library-management/appv1/router"
	"Library-management/appv1/tools"
	"fmt"
)

func Start() {
	defer func() {
		modle.Close()
	}()
	modle.New()
	tools.NewToken("")
	r := router.New()
	err := r.Run()
	if err != nil {
		fmt.Println("路由出错了：", err)
	}
}
