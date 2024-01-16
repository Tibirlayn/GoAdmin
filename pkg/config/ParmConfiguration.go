package config

import (
	"fmt"
	"log"
    "gorm.io/driver/sqlserver"  
    "gorm.io/gorm"
)

func ParmConfiguration() (*gorm.DB, error){
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	//подлючение к БД ...
    dns := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable", cfg.Parm.User, cfg.Parm.Password, cfg.Parm.Server, cfg.Parm.Port, cfg.Parm.DBname)
    db_parm, err := gorm.Open(sqlserver.Open(dns), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // Получаем объект базы данных gorm.DB и отложенно закрываем его соединение
    dbSQL, err := db_parm.DB()
    if err != nil {
        return nil, err
    }

	// Использование конфигурации
	fmt.Println("Server:", cfg.Parm.Server)
	fmt.Println("User:", cfg.Parm.User)
	fmt.Println("Passeord:", cfg.Parm.Password)
    fmt.Println("Port:", cfg.Parm.Port)
    fmt.Println("FNLParm:", cfg.Parm.DBname)

    // Проверка подключения 
    err = dbSQL.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
    return db_parm, nil
}