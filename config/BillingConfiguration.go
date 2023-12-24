package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func BillingConfiguration() {
	cfg, err := BillingLoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return 
    }



	//подлючение к БД
    connStringBilling := fmt.Sprintf("server=%s;user id=%s;password=%s;billing=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLBilling)
    db_billing, err := sql.Open("sqlserver", connStringBilling)
    if err != nil {
        log.Fatal(err)
    }
    defer db_billing.Close()

	// Использование конфигурации
	fmt.Println("Server:", cfg.Server)
	fmt.Println("User:", cfg.User)
	fmt.Println("Passeord:", cfg.Password)
    fmt.Println("FNLBilling:", cfg.FNLBilling)

    err = db_billing.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
}