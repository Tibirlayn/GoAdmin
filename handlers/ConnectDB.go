package handlers

import 	(
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
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

    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", cfg.Server, cfg.User, cfg.Password, cfg.Database)
    db, err := sql.Open("sqlserver", connString)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}