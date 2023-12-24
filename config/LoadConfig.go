package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AccountConfig struct {
	Server     string
	User       string
	Password   string
	FNLAccount string
}

type BattleConfig struct {
	Server    string
	User      string
	Password  string
	FNLBattle string
}

type BillingConfig struct {
	Server     string
	User       string
	Password   string
	FNLBilling string
}

type GameConfig struct {
	Server   string
	User     string
	Password string
	FNLGame  string
}

type LogConfig struct {
	Server   string
	User     string
	Password string
	FNLLog   string
}

type ParmConfig struct {
	Server   string
	User     string
	Password string
	FNLParm  string
}

type StatisticsConfig struct {
	Server        string
	User          string
	Password      string
	FNLStatistics string
}

func AccountLoadConfig() (AccountConfig, error) {
	var config AccountConfig

	// открываем и читаем файл, если файл нет, выдает ошибку
	viper.SetConfigFile("config/dbparam/accountConfig.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config, err
	}

	// Создаем структуру данных
	config.Server = viper.GetString("server")
	config.User = viper.GetString("user")
	config.Password = viper.GetString("password")
	config.FNLAccount = viper.GetString("account")

	//возращаем данные
	return config, nil
}

func BattleLoadConfig() (BattleConfig, error) {
	var config BattleConfig

	// открываем и читаем файл, если файл нет, выдает ошибку
	viper.SetConfigFile("config/dbparam/battleConfig.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config, err
	}

	// Создаем структуру данных
	config.Server = viper.GetString("server")
	config.User = viper.GetString("user")
	config.Password = viper.GetString("password")
	config.FNLBattle = viper.GetString("battle")

	//возращаем данные
	return config, nil
}

func BillingLoadConfig() (BillingConfig, error) {
	var config BillingConfig

	// открываем и читаем файл, если файл нет, выдает ошибку
	viper.SetConfigFile("config/dbparam/billingConfig.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config, err
	}

	// Создаем структуру данных
	config.Server = viper.GetString("server")
	config.User = viper.GetString("user")
	config.Password = viper.GetString("password")
	config.FNLBilling = viper.GetString("billing")

	//возращаем данные
	return config, nil
}

func GameLoadConfig() (GameConfig, error) {
	var config GameConfig

	// открываем и читаем файл, если файл нет, выдает ошибку
	viper.SetConfigFile("config/dbparam/gameConfig.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config, err
	}

	// Создаем структуру данных
	config.Server = viper.GetString("server")
	config.User = viper.GetString("user")
	config.Password = viper.GetString("password")
	config.FNLGame = viper.GetString("game")

	//возращаем данные
	return config, nil
}

func LogLoadConfig() (LogConfig, error) {
	var config LogConfig
	// открываем и читаем файл, если файл нет, выдает ошибку
	viper.SetConfigFile("config/dbparam/logConfig.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config, err
	}

	// Создаем структуру данных
	config.Server = viper.GetString("server")
	config.User = viper.GetString("user")
	config.Password = viper.GetString("password")
	config.FNLLog = viper.GetString("logs")

	//возращаем данные
	return config, nil
}

func ParmLoadConfig() (ParmConfig, error) {
	var config ParmConfig

	// открываем и читаем файл, если файл нет, выдает ошибку
	viper.SetConfigFile("config/dbparam/parmConfig.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config, err
	}

	// Создаем структуру данных
	config.Server = viper.GetString("server")
	config.User = viper.GetString("user")
	config.Password = viper.GetString("password")
	config.FNLParm = viper.GetString("parm")

	//возращаем данные
	return config, nil
}

func StatisticsLoadConfig() (StatisticsConfig, error) {
	var config StatisticsConfig

	// открываем и читаем файл, если файл нет, выдает ошибку
	viper.SetConfigFile("config/dbparam/statisticsConfig.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config, err
	}

	// Создаем структуру данных
	config.Server = viper.GetString("server")
	config.User = viper.GetString("user")
	config.Password = viper.GetString("password")
	config.FNLStatistics = viper.GetString("statistics")

	//возращаем данные
	return config, nil
}
