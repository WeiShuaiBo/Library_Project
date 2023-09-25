package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"library/appV5/common"
	//这里少了一句
)

var GlobalConn *gorm.DB

func New() {

	mysqlHost := common.V.GetString("mysql.host")
	mysqlPort := common.V.GetString("mysql.port")
	mysqlUsername := common.V.GetString("mysql.username")
	mysqlPassword := common.V.GetString("mysql.password")
	mysqldatavase := common.V.GetString("mysql.database")
	//parseTime=True&loc=Local MySQL 默认时间是格林尼治时间，与我们差八小时，需要定位到我们当地时间
	//my := mysqlHost

	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlUsername, mysqlPassword, mysqlHost+":"+mysqlPort, mysqldatavase)
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
