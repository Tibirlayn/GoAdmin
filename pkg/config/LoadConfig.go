package config

import (
	"fmt"

	"github.com/spf13/viper"
	//"github.com/Tibirlayn/GoAdmin/pkg/config/dbparam"
)

type Config struct {
	Account    DBConfig
	Battle     DBConfig
	Billing    DBConfig
	Game       DBConfig
	Log        DBConfig
	Parm       DBConfig
	Statistics DBConfig
}

type DBConfig struct {
	Server   string
	User     string
	Password string
	DBname   string
}

func LoadConfig() (Config, error) {
	var configFile Config
	// открываем и читаем файл, если файл нет, выдает ошибку
	viper.SetConfigFile("pkg/config/dbparam/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return configFile, err
	}

	// Записываем данные в структуру
	configData := Config{
		Account: DBConfig{
			Server:   viper.GetString("account.server"),
			User:     viper.GetString("account.user"),
			Password: viper.GetString("account.password"),
			DBname:   viper.GetString("account.account"),
		},
		Battle: DBConfig{
			Server:   viper.GetString("battle.server"),
			User:     viper.GetString("battle.user"),
			Password: viper.GetString("battle.password"),
			DBname:   viper.GetString("battle.battle"),
		},
		Billing: DBConfig{
			Server:   viper.GetString("billing.server"),
			User:     viper.GetString("billing.user"),
			Password: viper.GetString("billing.password"),
			DBname:   viper.GetString("billing.billing"),
		},
		Game: DBConfig{
			Server:   viper.GetString("game.server"),
			User:     viper.GetString("game.user"),
			Password: viper.GetString("game.password"),
			DBname:   viper.GetString("game.game"),
		},
		Log: DBConfig{
			Server:   viper.GetString("logs.server"),
			User:     viper.GetString("logs.user"),
			Password: viper.GetString("logs.password"),
			DBname:   viper.GetString("logs.logs"),
		},
		Parm: DBConfig{
			Server:   viper.GetString("parm.server"),
			User:     viper.GetString("parm.user"),
			Password: viper.GetString("parm.password"),
			DBname:   viper.GetString("parm.parm"),
		},
		Statistics: DBConfig{
			Server:   viper.GetString("statistics.server"),
			User:     viper.GetString("statistics.user"),
			Password: viper.GetString("statistics.password"),
			DBname:   viper.GetString("statistics.statistics"),
		},
	}

	return configData, nil
}
