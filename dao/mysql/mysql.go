package mysql

import (
	"Library_Project/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(cfg *config.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//设置了日志记录器的选项，使用默认的， 将日志模式设置为静默模式，不输出任何日志信息
		Logger: logger.Default.LogMode(logger.Silent),
		//设置了是否禁用自动ping功能，false表示不禁用自动ping功能，即在连接空闲时自动发送ping以保持连接活跃
		DisableAutomaticPing: false,
	})
	if err != nil {
		fmt.Println("数据库启动失败", err)
		return err
	}
	DB = db
	fmt.Println("数据库开启成功")
	return nil
}

// Close 关闭数据库连接
func Close() {
	sql, _ := DB.DB()
	err := sql.Close()
	if err != nil {
		fmt.Println("数据库关闭失败", err)
		return
	}
}
