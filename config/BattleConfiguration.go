package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func BattleConfiguration() (*sql.DB, error) {
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

    connStringBattle := fmt.Sprintf(
        "server=%s;user id=%s;password=%s;battle=%s;encrypt=disable", 
        cfg.Battle.Server, cfg.Battle.User, cfg.Battle.Password, cfg.Battle.DBname)
    db_battle, err := sql.Open("sqlserver", connStringBattle)
    if err != nil {
        log.Fatal(err)
    }
    defer db_battle.Close()
    
	// Использование конфигурации
	fmt.Println("Server:", cfg.Battle.Server)
	fmt.Println("User:", cfg.Battle.User)
	fmt.Println("Passeord:", cfg.Battle.Password)
    fmt.Println("FNLBattle:", cfg.Battle.DBname)

    // Проверка подключения 
    err = db_battle.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }

    return db_battle, nil
}