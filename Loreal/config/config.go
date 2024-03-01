package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     int
}

var ENV *Config

func LoadConfig() {
	log.Println("Load Server Configuration")
	log.Println(&ENV)
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
