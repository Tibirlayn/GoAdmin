package config

import (
	"fmt"
	"log"
    "gorm.io/driver/sqlserver"  
    "gorm.io/gorm"
)

func LogConfiguration() (*gorm.DB, error) {
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	//подлючение к БД ...
    dns := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable", cfg.Log.User, cfg.Log.Password, cfg.Log.Server, cfg.Log.Port, cfg.Log.DBname)
    db_log, err := gorm.Open(sqlserver.Open(dns), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // Получаем объект базы данных gorm.DB и отложенно закрываем его соединение
    dbSQL, err := db_log.DB()
    if err != nil {
        return nil, err
    }

	// Использование конфигурации
    fmt.Println("Server:", cfg.Log.Server)
	fmt.Println("User:", cfg.Log.User)
	fmt.Println("Passeord:", cfg.Log.Password)
    fmt.Println("Port:", cfg.Log.Port)
    fmt.Println("FNLParm:", cfg.Log.DBname)

    // Проверка подключения
    err = dbSQL.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
    return db_log, nil
}