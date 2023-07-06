package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"library/appV0/config"
)

var GlobalConn *gorm.DB
var RedisConn *redis.Pool

// Mysql 连接mysql
func Mysql(cfg *config.Mysql) error {
	//parseTime=True&loc=Local MySQL 默认时间是格林尼治时间，与我们差八小时，需要定位到我们当地时间
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.IP, cfg.Name)
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:%s\n", err)
		panic(err)
	}
	GlobalConn = conn
	return err
}

// Redis 连接redis
func Redis(cfg *config.Redis) error {
	redisCoon := &redis.Pool{
		MaxIdle:   cfg.MaxIdle,
		MaxActive: cfg.MaxActive,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.Ip)
		},
	}
	RedisConn = redisCoon
	return nil
}

func Close() {
	get := RedisConn.Get()
	db, _ := GlobalConn.DB()
	_ = get.Close()
	_ = db.Close()
}
