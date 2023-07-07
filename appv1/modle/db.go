package modle

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接数据库

var GlobalConn *gorm.DB

func New() {
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "127.0.0.1:3306", "library")
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
