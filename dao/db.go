package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalConn *gorm.DB

func New() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=True", "root", "123456", "127.0.0.1", 3306, "library")
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("数据库连接失败！")
	}
	GlobalConn = db
	_ = GlobalConn.AutoMigrate(Xxx{})
}
func Close() {
	db, _ := GlobalConn.DB()
	_ = db.Close()
}

//CREATE TABLE user(
//id int not null auto_increment,
//name varchar(40) default null,
//pwd varchar(40) default null,
//type bool default null,
//create_time datetime default null,
//primary key (`id`)
//) engine=InnoDB auto_increment=2 default charset=utf8mb4 collate=utf8mb4_bin;
//create table book(
//id bigint not null,
//book_name varchar(40),
//count int default null,
//create_time datetime default null,
//primary key (`id`)
//) engine=InnoDB auto_increment=2 default charset=utf8mb4 collate utf8mb4_bin;
//create table user_book(
//id int not null ,
//user_id int not null,
//book_id int not null,
//type bool not null,
//borrow_time datetime not null,
//return_time datetime default null,
//primary key (`id`)
//) engine=InnoDB auto_increment=2 default charset=utf8mb4 collate utf8mb4_bin;
