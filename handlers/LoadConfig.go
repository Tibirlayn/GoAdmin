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

	// открываем и читаем файл, если файл нет, выдает ошибку
	viper.SetConfigFile("config/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config, err
	}

	// Создаем структуру данных
	config.Server = viper.GetString("server")
	config.User = viper.GetString("user")
	config.Password = viper.GetString("password")
	config.Database = viper.GetString("database")

	//возращаем данные 
	return config, nil
}