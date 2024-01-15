package config

import (
	"fmt"
	"log"
    "gorm.io/driver/sqlserver"  
    "gorm.io/gorm"
    "github.com/Tibirlayn/GoAdmin/pkg/models/account"
)

func AccountConfiguration() (*gorm.DB, error) {
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }
	//подлючение к БД ...
    dns := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable", cfg.Account.User, cfg.Account.Password, cfg.Account.Server, cfg.Account.Port, cfg.Account.DBname)
    db_account, err := gorm.Open(sqlserver.Open(dns), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    // Автоматическое создание таблиц в базе данных MSSQL на основе определения модели User
    if err := db_account.AutoMigrate(&account.User{}); err != nil {
        panic(err)
    } 
    // Получаем объект базы данных gorm.DB и отложенно закрываем его соединение
    dbSQL, err := db_account.DB()
    if err != nil {
        return nil, err
    }
    // Использование конфигурации
	fmt.Println("Server:", cfg.Account.Server)
	fmt.Println("User:", cfg.Account.User)
	fmt.Println("Passeord:", cfg.Account.Password)
    fmt.Println("Port:", cfg.Account.Port)
	fmt.Println("FNLAccount:", cfg.Account.DBname)

	if err := dbSQL.Ping(); err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
        return nil, err
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }

    return db_account, nil
}

