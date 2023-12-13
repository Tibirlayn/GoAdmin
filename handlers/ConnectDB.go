package handlers

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectDB() {

	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return 
    }

	// Использование конфигурации
	fmt.Println("Server:", cfg.Server)
	fmt.Println("Port:", cfg.User)
	fmt.Println("Debug:", cfg.Password)
	fmt.Println("Debug:", cfg.Database)

	//подлючение к БД
    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.Database)
    db, err := sql.Open("sqlserver", connString)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	err = db.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
    } else {
        fmt.Println("Успешное подключение к базе данных")
    }

}