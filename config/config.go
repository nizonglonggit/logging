package config

import "github.com/BurntSushi/toml"

var (
	Conf *Config
)

// Config 对应配置文件结构
type Config struct {
	Lumberjacks map[string]Lumberjack `toml:"lumberjacks"`
}

// UnmarshalConfig 解析toml配置
func UnmarshalConfig(tomlFile string) (*Config, error) {
	if _, err := toml.DecodeFile(tomlFile, &Conf); err != nil {
		return Conf, err
	}

	return Conf, nil
}

type Lumberjack struct {
	MaxSize    int  `json:"max_size"`    //最大M数，超过则切割
	MaxBackups int  `json:"max_backups"` //最大文件保留数，超过就删除最老的日志文件
	MaxAge     int  `json:"max_age"`     //保存30天
	Compress   bool `json:"compress"`    //是否压缩
}

func init() {
	// 解析配置文件
	_, err := UnmarshalConfig("../docs/log.toml")
	if err != nil {
		panic(err)
	}
}
