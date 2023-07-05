package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	Log.SetFormatter(&logrus.JSONFormatter{})

	Log.SetLevel(logrus.InfoLevel)

	//在文件中打印日志
	//file, err := os.OpenFile("library_app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//
	//if err == nil {
	//	Log.SetOutput(file)
	//} else {
	//	Log.Info("Failed to log to file,using default stderr")
	//	fmt.Printf("日志文件创建失败err:%s\n", err)
	//
	//}
	Log.SetOutput(os.Stdout)
	Log.Info("Logrous configured successfully!")
}
