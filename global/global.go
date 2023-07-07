package global

import (
	"Library_Project/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	FAST_CONFIG config.Server
	FAST_VP     *viper.Viper
	FAST_LOG    *zap.Logger
	FAST_DB     *gorm.DB
)
