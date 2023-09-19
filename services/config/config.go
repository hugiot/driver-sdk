package config

import (
	"github.com/spf13/viper"
	"log"
)

var defaultFileName = "config.toml"

func New() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(defaultFileName)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("init config error", err)
	}

	return v
}
