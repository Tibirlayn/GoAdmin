package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func StatisticsConfiguration() (*sql.DB, error){
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	//подлючение к БД
    connStringStatistics := fmt.Sprintf(
        "server=%s;user id=%s;password=%s;statistics=%s;encrypt=disable", 
        cfg.Statistics.Server, cfg.Statistics.User, cfg.Statistics.Password, cfg.Statistics.DBname)
    db_statistics, err := sql.Open("sqlserver", connStringStatistics)
    if err != nil {
        log.Fatal(err)
    }
    defer db_statistics.Close()

	// Использование конфигурации
	fmt.Println("Server:", cfg.Statistics.Server)
	fmt.Println("User:", cfg.Statistics.User)
	fmt.Println("Passeord:", cfg.Statistics.Password)
    fmt.Println("FNLStatistics:", cfg.Statistics.DBname)

    // Проверка подключения
    err = db_statistics.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
    return db_statistics, nil
}