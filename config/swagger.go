package config

type Swagger struct {
	Enabled     bool     `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Title       string   `mapstructure:"title" json:"title" yaml:"title"`
	Description string   `mapstructure:"description" json:"description" yaml:"description"`
	Version     string   `mapstructure:"version" json:"version" yaml:"version"`
	Host        string   `mapstructure:"host" json:"host" yaml:"host"`
	BasePath    string   `mapstructure:"base-path" json:"base-path" yaml:"base-path"`
	Schemes     []string `mapstructure:"schemes" json:"schemes" yaml:"schemes"`
}
