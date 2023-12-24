package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func GameConfiguration() {
	cfg, err := GameLoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return 
    }

	//подлючение к БД
    connStringGame := fmt.Sprintf("server=%s;user id=%s;password=%s;game=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLGame)
    db_game, err := sql.Open("sqlserver", connStringGame)
    if err != nil {
        log.Fatal(err)
    }
    defer db_game.Close()

	// Использование конфигурации
	fmt.Println("Server:", cfg.Server)
	fmt.Println("User:", cfg.User)
	fmt.Println("Passeord:", cfg.Password)
    fmt.Println("FNLGame:", cfg.FNLGame)

    err = db_game.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
}