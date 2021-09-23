package config

import (
	"github.com/Toolnado/SwalloW/pkg/utils"
	"github.com/spf13/viper"
)

const (
	typeFile = "yaml"
	nameFile = "config"
	pathFile = "./"
)

type Config struct {
	ServerPort       string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabaseName     string
	DatabaseSSLMode  string
	DatabasePassword string
}

func NewConfig() *Config {
	utils.GetConfig(typeFile, nameFile, pathFile)

	return &Config{
		ServerPort:       viper.GetString("port"),
		DatabaseHost:     viper.GetString("db.host"),
		DatabasePort:     viper.GetString("db.port"),
		DatabaseUser:     viper.GetString("db.user"),
		DatabaseName:     viper.GetString("db.name"),
		DatabaseSSLMode:  viper.GetString("db.sslmode"),
		DatabasePassword: viper.GetString("db.password"),
	}
}
