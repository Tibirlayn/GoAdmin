package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func ParmConfiguration() (*sql.DB, error){
	cfg, err := LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return nil, err
    }

	//подлючение к БД
    connStringParm := fmt.Sprintf(
        "server=%s;user id=%s;password=%s;parm=%s;encrypt=disable", 
        cfg.Parm.Server, cfg.Parm.User, cfg.Parm.Password, cfg.Parm.DBname)
    db_parm, err := sql.Open("sqlserver", connStringParm)
    if err != nil {
        log.Fatal(err)
    }
    defer db_parm.Close()

	// Использование конфигурации
	fmt.Println("Server:", cfg.Parm.Server)
	fmt.Println("User:", cfg.Parm.User)
	fmt.Println("Passeord:", cfg.Parm.Password)
    fmt.Println("FNLParm:", cfg.Parm.DBname)

    // Проверка подключения 
    err = db_parm.Ping()
	if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        fmt.Println("-----------------------------------------")
    } else {
        fmt.Println("Успешное подключение к базе данных")
        fmt.Println("-----------------------------------------")
    }
    return db_parm, nil
}