package config

import (
	"fmt"
	"log"
    "gorm.io/driver/sqlserver"  
    "gorm.io/gorm"
)

func BillingConfiguration() (*gorm.DB, error) {
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	//подлючение к БД ...
    dns := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable", cfg.Billing.User, cfg.Billing.Password, cfg.Billing.Server, cfg.Billing.Port, cfg.Billing.DBname)
    db_billing, err := gorm.Open(sqlserver.Open(dns), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    // Получаем объект базы данных gorm.DB и отложенно закрываем его соединение
    dbSQL, err := db_billing.DB()
    if err != nil {
        return nil, err
    }

	// Использование конфигурации
	fmt.Println("Server:", cfg.Billing.Server)
	fmt.Println("User:", cfg.Billing.User)
	fmt.Println("Passeord:", cfg.Billing.Password)
    fmt.Println("Port:", cfg.Billing.Port)
    fmt.Println("FNLBilling:", cfg.Billing.DBname)

    // Проверка подключения 
    err = dbSQL.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }

    return db_billing, err
}