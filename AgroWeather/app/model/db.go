package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalConn *gorm.DB

// New 连接数据库
func New() {
	//parseTime=True&loc=Local MySQL 默认时间是格林尼治时间，与我们差八小时，需要定位到我们当地时间
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "020409", "127.0.0.1:3306", "agroweather")
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:%s\n", err)
		panic(err)
	}
	GlobalConn = conn
}

func Close() {
	db, _ := GlobalConn.DB()
	_ = db.Close()
}
