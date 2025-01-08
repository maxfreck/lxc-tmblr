package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Container struct {
	Root         string
	Dependencies []string
}

type AppConfig struct {
	Hello      string
	Containers map[string]Container
}

func GetAppConfig() *AppConfig {
	var Config AppConfig

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/lxc-tmblr/")
	viper.AddConfigPath("$HOME/.config/lxc-tmblr/")
	viper.AddConfigPath(".")

	err1 := viper.ReadInConfig()
	if err1 != nil {
		panic(fmt.Errorf("fatal error config file: %w", err1))
	}

	err2 := viper.Unmarshal(&Config)
	if err2 != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err2))
	}

	return &Config
}
