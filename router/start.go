package router

import (
	"Library_Project/config"
	"Library_Project/dao/mysql"
	"Library_Project/dao/redis"
	"Library_Project/utils/jwt"
	"Library_Project/utils/logger"
	"Library_Project/utils/snowflake"
	"fmt"
	"log"
)

func Start() {
	if err := Init(); err != nil {
		fmt.Printf("配置加载失败了，err:%v", err)
		return
	}
	defer func() {
		mysql.Close()
		redis.Close()
	}()
	//重新定义秘钥
	jwt.NewToken("Go")

	r := Router()
	err := r.Run(":8083")
	if err != nil {
		log.Fatalf("路由启动失败，请检查")
		return
	}
}

// 基础配置配置
func Init() (err error) {
	if err := config.Init(); err != nil {
		fmt.Printf("load config failed,err:%v\n", err)
		return
	}
	if err1 := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return err1
	}
	if err2 := mysql.Init(config.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return err2
	}
	if err3 := redis.Init(config.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err:%v\n", err)
		return err3
	}
	if err4 := snowflake.Init(1); err4 != nil {
		return err4
	}
	return nil
}
