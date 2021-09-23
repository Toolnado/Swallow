package utils

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func GetConfig(typeFile, nameFile, pathFile string) {
	viper.SetConfigType(typeFile)
	viper.SetConfigName(nameFile)
	viper.AddConfigPath(pathFile)
	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("[fatal error config file: %s ]\n", err)
	}
}

func GetEnvStr(name string, def string) string {
	value := os.Getenv(name)
	if len(value) == 0 {
		return def
	}
	return value
}
