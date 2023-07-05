package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"library/logger"
)

var GlobalConn *gorm.DB

func New() {
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "313313", "127.0.0.1:3306", "library")
	conn, err := gorm.Open("mysql", my)
	if err != nil {
		logger.Log.Error("数据库连接失败", err)
		panic(err)
	}
	GlobalConn = conn

	GlobalConn.AutoMigrate(&User{}, &Admin{}, &Book{}, &LendBooks{})
	logger.Log.Info("数据库连接成功")
}
func Close() {
	_ = GlobalConn.Close()
}
