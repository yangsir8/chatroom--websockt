package config

import (
	_"fmt"
	"github.com/spf13/viper"
)


type Config struct {
	Filename string
}

//读取配置
func (c *Config) InitConfig() error {
	if c.Filename != "" {
		viper.SetConfigFile(c.Filename)
	}else{
		viper.AddConfigPath("config")
		viper.SetConfigName("conf")
	}
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}

