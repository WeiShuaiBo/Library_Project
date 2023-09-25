package common

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var V *viper.Viper

func InitConfig() {
	workDir, _ := os.Getwd()
	V = viper.New()
	V.SetConfigName("conf")
	V.SetConfigType("yaml")
	V.AddConfigPath(workDir + "/appV5/config/")
	err := V.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Load Config Error: %s", err.Error()))
	}
	fmt.Println(viper.GetString("server.port"))
}
