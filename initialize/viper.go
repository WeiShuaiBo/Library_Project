package initialize

import (
	"Library_Project/global"
	"Library_Project/utils"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Viper(path ...string) *viper.Viper {

	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()

		if config == "" {
			// 优先级: main > 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
			} else {
				config = configEnv
				fmt.Printf("您正在使用的config是名为'FAST_CONFIG'的环境变量,configde的路径为%v\n", config)
			}

		} else {
			fmt.Printf("您正在使用命令行-c参数传递的配置，config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
		if err := v.Unmarshal(&global.FAST_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.FAST_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v

}
