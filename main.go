package main

import (
	//"database/sql"
	"fmt"
	//"net"
	_ "github.com/denisenkom/go-mssqldb"
	//"config/viperHandle"
	//"log"
	"github.com/Tibirlayn/GoAdmin/handlers"
)
func main() {

	cfg, err := handlers.LoadConfig()
    if err != nil {
        fmt.Println("Error loading config:", err)
        return
    }

	    // Использование конфигурации
		fmt.Println("Server:", cfg.Server)
		fmt.Println("Port:", cfg.User)
		fmt.Println("Debug:", cfg.Password)
		fmt.Println("Debug:", cfg.Database)

	//viper.SetConfigFile("config/config.yml")
	//err := viper.ReadInConfig()
	//if err != nil {
	//	fmt.Println("Error reading config file:", err)
	//}
//
	//server := viper.GetString("server")
	//user := viper.GetString("user")
	//password := viper.GetString("password")
	//database := viper.GetString("database")
//
	//fmt.Println("Server ", server)
	//fmt.Println("User ", user)
	//fmt.Println("Password ", password)
	//fmt.Println("Database ", database)

	//server := "your-server"
    //user := "your-user"
    //password := "your-password"
    //database := "your-database"
//
    //connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", server, user, password, database)
    //db, err := sql.Open("sqlserver", connString)
    //if err != nil {
    //    log.Fatal(err)
    //}
    //defer db.Close()
}