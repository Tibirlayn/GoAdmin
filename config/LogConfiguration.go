package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func LogConfiguration() {
	cfg, err := LogLoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return 
    }

	//подлючение к БД
    connStringLog := fmt.Sprintf("server=%s;user id=%s;password=%s;logs=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLLog)
    db_log, err := sql.Open("sqlserver", connStringLog)
    if err != nil {
        log.Fatal(err)
    }
    defer db_log.Close()

	// Использование конфигурации
	fmt.Println("Server:", cfg.Server)
	fmt.Println("User:", cfg.User)
	fmt.Println("Passeord:", cfg.Password)
    fmt.Println("FNLLog:", cfg.FNLLog)

    err = db_log.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
}