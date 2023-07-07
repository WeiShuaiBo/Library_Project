package initialize

import (
	"Library_Project/global"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func GormMysql01() *gorm.DB {

	msq := global.FAST_CONFIG.Mysql
	if msq.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       msq.Dsn(), // DSN data source name
		DefaultStringSize:         191,       // string 类型字段的默认长度
		SkipInitializeWithVersion: false,     // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(msq.MaxIdleConns)
		sqlDB.SetMaxOpenConns(msq.MaxOpenConns)
		return db
	}
	return nil
}

func gormConfig() *gorm.Config {
	config := &gorm.Config{
		//DisableForeignKeyConstraintWhenMigrating: true,
		//表复数禁用
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	return config
}

func GormMysql() *gorm.DB {
	msq := global.FAST_CONFIG.Mysql
	fmt.Println(msq.Dsn())
	sqlDb, err := sql.Open("mysql", msq.Dsn())
	if err != nil {
		zap.L().Fatal("数据库连接出错：", zap.Error(err))
	}
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDb}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		zap.L().Fatal("数据库连接出错：", zap.Error(err))
	} else {
		zap.L().Info("数据库连接成功")
	}
	return db
}
