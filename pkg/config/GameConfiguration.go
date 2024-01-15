package config

import (
	"fmt"
	"log"
    "gorm.io/driver/sqlserver"  
    "gorm.io/gorm"
)

func GameConfiguration() (*gorm.DB, error){
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	//подлючение к БД ...
    dns := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable", cfg.Game.User, cfg.Game.Password, cfg.Game.Server, cfg.Game.Port, cfg.Game.DBname)
    db_game, err := gorm.Open(sqlserver.Open(dns), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // Получаем объект базы данных gorm.DB и отложенно закрываем его соединение
    dbSQL, err := db_game.DB()
    if err != nil {
        return nil, err
    }

	// Использование конфигурации
	fmt.Println("Server:", cfg.Game.Server)
	fmt.Println("User:", cfg.Game.User)
	fmt.Println("Passeord:", cfg.Game.Password)
    fmt.Println("Port:", cfg.Game.Port)
    fmt.Println("FNLGame:", cfg.Game.DBname)

    // Проверка подключения
    err = dbSQL.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
    return db_game, nil
}