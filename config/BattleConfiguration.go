package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func BattleConfiguration() {
	cfg, err := BattleLoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return 
    }

    connStringBattle := fmt.Sprintf("server=%s;user id=%s;password=%s;battle=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLBattle)
    db_battle, err := sql.Open("sqlserver", connStringBattle)
    if err != nil {
        log.Fatal(err)
    }
    defer db_battle.Close()
    
	// Использование конфигурации
	fmt.Println("Server:", cfg.Server)
	fmt.Println("User:", cfg.User)
	fmt.Println("Passeord:", cfg.Password)
    fmt.Println("FNLBattle:", cfg.FNLBattle)

    err = db_battle.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
}