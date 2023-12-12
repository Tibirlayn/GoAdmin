package handlers

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server string
	User string
	Password string
	Database string
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile("config/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config, err
	}

	config.Server = viper.GetString("server")
	config.User = viper.GetString("user")
	config.Password = viper.GetString("password")
	config.Database = viper.GetString("database")

	return config, nil
}