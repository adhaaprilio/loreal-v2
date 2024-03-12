package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB_NAME     string
	DB_PORT     int
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
}

var ENV *Config

func LoadConfig() {
	fmt.Println("Load Server Configuration")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&ENV)
	if err != nil {
		panic(err)
	}
}
