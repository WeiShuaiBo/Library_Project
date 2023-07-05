// @Author	zhangjiaozhu 2023/7/3 14:50:00
package main

import (
	"MyBook/common/config"
	"MyBook/common/snowflake"
	"MyBook/models"
	"MyBook/models/DB/MysqlDB"
	redis "MyBook/models/DB/RedisDB"
	"MyBook/routers"
	"fmt"
)

// @title			标题
// @version		1.0
// @description	描述信息
// @termsOfService	http://swagger.io/terms/
// @contact.name	联系人信息
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	// 加载配置文件
	config.NewConfig()
	// 加载雪花算法
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 加载redis
	err := redis.InitRedisDb()
	if err != nil {
		panic("redis加载失败")
	}
	// 加载mysql数据库
	MysqlDB.NewMySql()
	// 迁移表结构
	MysqlDB.DB.AutoMigrate(&models.User{}, &models.Book{}, &models.BorrowRecord{})
	// 加载路由
	routers.NewRouter()
}
