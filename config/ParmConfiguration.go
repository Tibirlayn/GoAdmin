package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func ParmConfiguration() {
		cfg, err := ParmLoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return 
    }

	// Использование конфигурации
	fmt.Println("Server:", cfg.Server)
	fmt.Println("User:", cfg.User)
	fmt.Println("Passeord:", cfg.Password)
    fmt.Println("FNLParm:", cfg.FNLParm)

	//подлючение к БД
    connStringParm := fmt.Sprintf("server=%s;user id=%s;password=%s;parm=%s;encrypt=disable", cfg.Server, cfg.User, cfg.Password, cfg.FNLParm)
    db_parm, err := sql.Open("sqlserver", connStringParm)
    if err != nil {
        log.Fatal(err)
    }
    defer db_parm.Close()

    err = db_parm.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
    } else {
        fmt.Println("Успешное подключение к базе данных")
    }
}