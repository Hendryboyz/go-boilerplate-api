package bootstrap

import (
	"fmt"
	"go-boilerplate-api/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfiguration() *viper.Viper {
	var cfgFile string
	if envConfigFile := os.Getenv("SERVER_CONFIG"); envConfigFile != "" {
		cfgFile = envConfigFile
	}
	config := viper.New()
	if cfgFile != "" {
		config.SetConfigFile(cfgFile)
	} else {
		config.SetConfigName("local")
		config.SetConfigType("yaml")

		config.AddConfigPath("../configs/")
		config.AddConfigPath("configs/")
	}

	config.AutomaticEnv()

	// find and read config
	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Fprintln(os.Stderr, "Using config file:", config.ConfigFileUsed())

	config.WatchConfig()
	config.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)

		loadConfigStruct(config)
	})

	loadConfigStruct(config)

	return config
}

func loadConfigStruct(config *viper.Viper) {
	if err := config.Unmarshal(&global.App.Config); err != nil {
		fmt.Println("Failed to unmarshal config file:", err)
	}
}
