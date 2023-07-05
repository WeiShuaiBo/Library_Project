package logger

import (
	"Library_Project/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var lg *zap.Logger

func Init(cfg *config.LogConfig, mode string) (err error) {
	//日志写入器：
	writeSyncer := getLogWriter(cfg.FileName, cfg.MaxSize, cfg.MaxBackUps, cfg.MaxAge)
	encode := getEncode()
	//将文本表示的日志级别解析成对应的zapcore.level类型
	//l是一个zapcore.Level类型的指针，他讲配置对象中的日志级别解析为相应的zapcore.level类型，unmarshalText方法用于将文本表示的日志级别转换为对应的枚举值
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if mode == "dev" {
		//进入开发模式，日志输出到终端
		conosoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encode, writeSyncer, l),
			zapcore.NewCore(conosoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {

		core = zapcore.NewCore(encode, writeSyncer, l)
	}
	//创建一个logger实例，并通过zap.AddCaller()方法添加调用者信息
	lg = zap.New(core, zap.AddCaller())
	//将全局日志记录器替换为新的logger
	zap.ReplaceGlobals(lg)
	//记录一条初始化成功的日志消息
	zap.L().Info("init logger success")
	return nil
}

func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}
	//将日志记录器logger 添加到zap日志核心 向日志核心添加同步的日志记录器
	//同步日志记录器：当日志记录器接收到日志的时候，会立即将其写入目标输出位置
	return zapcore.AddSync(lumberJackLogger)
}

// 将日志时间编码为字节流接口
func getEncode() zapcore.Encoder {
	//创建一个预定义的生产环境下的编码器配置对象
	encoderConfig := zap.NewProductionEncoderConfig()
	//规定编码格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//日志中的时间字段的键名是 time
	encoderConfig.TimeKey = "time"
	// 指定日志级别的编码格式为大写形式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//持续时间的编码格式为秒
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	//调用方的编码格式为短路径形式  只包含文件名和行号
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
