package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func GameConfiguration() (*sql.DB, error){
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	//подлючение к БД
    connStringGame := fmt.Sprintf(
        "server=%s;user id=%s;password=%s;game=%s;encrypt=disable", 
        cfg.Game.Server, cfg.Game.User, cfg.Game.Password, cfg.Game.DBname)
    db_game, err := sql.Open("sqlserver", connStringGame)
    if err != nil {
        log.Fatal(err)
    }
    defer db_game.Close()

	// Использование конфигурации
	fmt.Println("Server:", cfg.Game.Server)
	fmt.Println("User:", cfg.Game.User)
	fmt.Println("Passeord:", cfg.Game.Password)
    fmt.Println("FNLGame:", cfg.Game.DBname)

    // Проверка подключения
    err = db_game.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
    return db_game, nil
}