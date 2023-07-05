// @Author	zhangjiaozhu 2023/7/3 15:22:00
package config

import (
	"os"

	"github.com/spf13/viper"
)

var Config *Conf

type Conf struct {
	System *System           `yaml:"system"`
	Mysql  map[string]*MySql `yaml:"mysql"`
	Redis  *Redis            `yaml:"redis"`
}
type System struct {
	AppEnv      string `yaml:"appEnv"`
	Domain      string `yaml:"domain"`
	Version     string `yaml:"version"`
	HttpPort    string `yaml:"httpPort"`
	Host        string `yaml:"host"`
	UploadModel string `yaml:"uploadModel"`
}

type MySql struct {
	Dialect  string `yaml:"dialect"`
	DbHost   string `yaml:"dbHost"`
	DbPort   string `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}
type Redis struct {
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisUsername string `yaml:"redisUsername"`
	RedisPassword string `yaml:"redisPwd"`
	RedisDbName   int    `yaml:"redisDbName"`
	RedisNetwork  string `yaml:"redisNetwork"`
}

func NewConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/")
	if err := viper.ReadInConfig(); err != nil {
		panic("找不到配置文件")
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic("解析配置文件失败")
	}
}
