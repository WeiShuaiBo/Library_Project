package config

type System struct {
	Env               string `mapstructure:"env" json:"env" yaml:"env"`                                          // 环境值
	Port              int    `mapstructure:"port" json:"port" yaml:"port"`                                       // 端口值
	EnabledMultipoint bool   `mapstructure:"enabled-multipoint" json:"enabled-multipoint" yaml:"use-multipoint"` // 是否允许多端登录
	OssType           string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                            // Oss类型
	EnableCors        bool   `mapstructure:"enable-cors" json:"enable-cors" yaml:"enable-cors"`                  // 是否允许跨域
	EnableRedis       bool   `mapstructure:"enable-redis" json:"enable-redis" yaml:"enable-redis"`               //是否允许redis支持
}
