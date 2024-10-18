package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(environment string) *viper.Viper {
	config = viper.New()

	config.SetConfigName(environment)
	config.SetConfigType("yaml")

	// allow multiple config search paths
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")

	// find and read config
	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return config
}

func GetConfig() *viper.Viper {
	return config
}
