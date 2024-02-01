package config

import (
	"fmt"
	"log"
    "gorm.io/driver/sqlserver"  
    "gorm.io/gorm"
)

func StatisticsConfiguration() (*gorm.DB, error){
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

    //подлючение к БД ...
    dns := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable", cfg.Statistics.User, cfg.Statistics.Password, cfg.Statistics.Server, cfg.Statistics.Port, cfg.Statistics.DBname)
    db_statistics, err := gorm.Open(sqlserver.Open(dns), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // Получаем объект базы данных gorm.DB и отложенно закрываем его соединение
    dbSQL, err := db_statistics.DB()
    if err != nil {
        return nil, err
    }

    // Использование конфигурации
    fmt.Println("Server:", cfg.Statistics.Server)
    fmt.Println("User:", cfg.Statistics.User)
    fmt.Println("Passeord:", cfg.Statistics.Password)
    fmt.Println("Port:", cfg.Statistics.Port)
    fmt.Println("FNLParm:", cfg.Statistics.DBname)

    // Проверка подключения
    err = dbSQL.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
    return db_statistics, nil
}