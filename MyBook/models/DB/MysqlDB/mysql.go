// @Author	zhangjiaozhu 2023/7/3 15:56:00
package MysqlDB

import (
	"MyBook/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var DB *gorm.DB

func NewMySql() {
	mConfig := config.Config.Mysql["default"]
	// mysql 读 主
	pathRead := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":", mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       pathRead, // DSN 数据库的名字
		DefaultStringSize:         256,      // string类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("链接数据库失败")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池 ， 空闲
	sqlDB.SetMaxOpenConns(100) // 连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
}
