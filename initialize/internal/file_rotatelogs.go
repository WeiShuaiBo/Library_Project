package internal

import (
	"Library_Project/global"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

// 日志切割归档
var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotatelogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	// New() 创建一个循环日志对象
	fileWriter, err := rotatelogs.New(

		path.Join(global.FAST_CONFIG.Zap.Director, "%Y-%m-%d", level+".log"),
		// WithClock() 获取当前时间
		rotatelogs.WithClock(rotatelogs.Local),
		// 日志留存时间
		rotatelogs.WithMaxAge(time.Duration(global.FAST_CONFIG.Zap.MaxAge)*24*time.Hour),
		// 切割时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	//fmt.Println(path.Join(global.FAST_CONFIG.Zap.Director))
	if global.FAST_CONFIG.Zap.LogInConsole {
		// NewMultiWriteSyncer() 将日志写入多个地方。
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
