package appV0

import (
	"fmt"
	"library/appV0/config"
	"library/appV0/logger"
	"library/appV0/model"
	"library/appV0/router"
	"library/appV0/tools"
	"os"
)

func Start() {
	defer func() {
		model.Close()
	}()

	// load config from config.json
	if len(os.Args) < 1 {
		return
	}

	if err := config.Init("./config.json"); err != nil {
		panic(err)
	}

	// init logger
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// init mysql
	if err := model.Mysql(config.Conf.Mysql); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	// init redis
	if err := model.Redis(config.Conf.Redis); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}

	tools.NewToken("")

	r := router.New()

	addr := fmt.Sprintf(":%v", config.Conf.Port)

	_ = r.Run(addr)
}
