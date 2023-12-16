package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func AccountConfiguration() {
    
	cfg, err := AccountLoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return 
    }

	// Использование конфигурации
	fmt.Println("Server:", cfg.Server)
	fmt.Println("User:", cfg.User)
	fmt.Println("Passeord:", cfg.Password)
	fmt.Println("FNLAccount:", cfg.FNLAccount)

	//подлючение к БД
    connStringAccount := fmt.Sprintf("server=%s;user id=%s;password=%s;account=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLAccount)
    db_account, err := sql.Open("sqlserver", connStringAccount)
    if err != nil {
        log.Fatal(err)
    }
    defer db_account.Close()

	err = db_account.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
    } else {
        fmt.Println("Успешное подключение к базе данных")
    }
}