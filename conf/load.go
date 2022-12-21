package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()
	// 读取 toml 格式配置文件
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path: %s, %s", filePath, err)
	}

	return nil
}

// 从环境变量加载配置
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return err
	}

	return nil
}
