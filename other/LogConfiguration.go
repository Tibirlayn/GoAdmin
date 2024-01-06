package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func LogConfiguration() (*sql.DB, error) {
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	//подлючение к БД
    connStringLog := fmt.Sprintf("server=%s;user id=%s;password=%s;logs=%s;encrypt=disable", cfg.Log.Server, cfg.Log.User, cfg.Log.Password, cfg.Log.DBname)
    db_log, err := sql.Open("sqlserver", connStringLog)
    if err != nil {
        log.Fatal(err)
    }
    defer db_log.Close()

	// Использование конфигурации
	fmt.Println("Server:", cfg.Log.Server)
	fmt.Println("User:", cfg.Log.User)
	fmt.Println("Passeord:", cfg.Log.Password)
    fmt.Println("FNLLog:", cfg.Log.DBname)

    // Проверка подключения
    err = db_log.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
    return db_log, nil
}