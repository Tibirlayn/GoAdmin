package account

import (
	"fmt"
	"github.com/Tibirlayn/GoAdmin/handlers"
)

func User() {

	// подключение к БД
	db := handlers.ConnectDB()

	rows, err := db.Query("SELECT * FROM TblUser LIMIT 20")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		// Сканирование значений из результата запроса
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}
}