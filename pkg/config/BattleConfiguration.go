package config

import (
	"fmt"
	"log"
    "gorm.io/driver/sqlserver"  
    "gorm.io/gorm"
)

func BattleConfiguration() (*gorm.DB, error) {
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	//подлючение к БД ...
    dns := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable", cfg.Battle.User, cfg.Battle.Password, cfg.Battle.Server, cfg.Battle.Port, cfg.Battle.DBname)
    db_battle, err := gorm.Open(sqlserver.Open(dns), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // Получаем объект базы данных gorm.DB и отложенно закрываем его соединение
    dbSQL, err := db_battle.DB()
    if err != nil {
        return nil, err
    }

	// Использование конфигурации
	fmt.Println("Server:", cfg.Battle.Server)
	fmt.Println("User:", cfg.Battle.User)
	fmt.Println("Passeord:", cfg.Battle.Password)
    fmt.Println("Port:", cfg.Battle.Port)
    fmt.Println("FNLParm:", cfg.Battle.DBname)

    // Проверка подключения 
    err = dbSQL.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
    return db_battle, nil
}