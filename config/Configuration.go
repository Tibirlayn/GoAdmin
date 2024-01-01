package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func Configuration() (*sql.DB, error) {
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	// Подлючение к БД
    connStringAccount := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;account=%s;encrypt=disable", 
		cfg.Account.Server, cfg.Account.User, cfg.Account.Password, cfg.Account.DBname)
    db_account, err := sql.Open("sqlserver", connStringAccount)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    defer db_account.Close()

    // Использование конфигурации
	fmt.Println("Server:", cfg.Account.Server)
	fmt.Println("User:", cfg.Account.User)
	fmt.Println("Passeord:", cfg.Account.Password)
	fmt.Println("FNLAccount:", cfg.Account.DBname)

	// Проверка подключения 
	err = db_account.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
        return nil, err
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }

    return db_account, nil
}