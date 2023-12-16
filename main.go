package main

import (
	"github.com/Tibirlayn/GoAdmin/handlers"
	"github.com/gin-gonic/gin"
)
func main() {
	// Создаем новый маршрутизатор Gin
	router := gin.Default()
	


	//	* узнать данные бд
	//	* работа с данными из бд

	// Обработчик GET/POST/PUT запроса на корневой URL
	handlers.Routers(router)

	// Запустить сервер на порту 8080
	router.Run(":8080")

}
