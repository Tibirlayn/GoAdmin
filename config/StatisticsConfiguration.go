package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func StatisticsConfiguration() {
	cfg, err := StatisticsLoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return 
    }

	//подлючение к БД
    connStringStatistics := fmt.Sprintf("server=%s;user id=%s;password=%s;statistics=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLStatistics)
    db_statistics, err := sql.Open("sqlserver", connStringStatistics)
    if err != nil {
        log.Fatal(err)
    }
    defer db_statistics.Close()

	// Использование конфигурации
	fmt.Println("Server:", cfg.Server)
	fmt.Println("User:", cfg.User)
	fmt.Println("Passeord:", cfg.Password)
    fmt.Println("FNLStatistics:", cfg.FNLStatistics)

    err = db_statistics.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
}