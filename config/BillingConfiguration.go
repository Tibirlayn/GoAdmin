package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func BillingConfiguration() (*sql.DB, error) {
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	//подлючение к БД
    connStringBilling := fmt.Sprintf(
        "server=%s;user id=%s;password=%s;billing=%s;encrypt=disable", 
    cfg.Billing.Server, cfg.Billing.User, cfg.Billing.Password, cfg.Billing.DBname)
    db_billing, err := sql.Open("sqlserver", connStringBilling)
    if err != nil {
        log.Fatal(err)
    }
    defer db_billing.Close()

	// Использование конфигурации
	fmt.Println("Server:", cfg.Billing.Server)
	fmt.Println("User:", cfg.Billing.User)
	fmt.Println("Passeord:", cfg.Billing.Password)
    fmt.Println("FNLBilling:", cfg.Billing.DBname)

    // Проверка подключения 
    err = db_billing.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }

    return nil, err
}