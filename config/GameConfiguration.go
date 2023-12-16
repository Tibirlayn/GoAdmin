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

	// Использование конфигурации
	fmt.Println("Server:", cfg.Server)
	fmt.Println("User:", cfg.User)
	fmt.Println("Passeord:", cfg.Password)
    fmt.Println("FNLGame2155:", cfg.FNLGame)

	//подлючение к БД
    connStringGame2155 := fmt.Sprintf("server=%s;user id=%s;password=%s;game2155=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLGame)
    db_game2155, err := sql.Open("sqlserver", connStringGame2155)
    if err != nil {
        log.Fatal(err)
    }
    defer db_game2155.Close()


    err = db_game2155.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
    } else {
        fmt.Println("Успешное подключение к базе данных")
    }
}