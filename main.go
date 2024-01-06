package main

import (
	"log"
	"github.com/Tibirlayn/GoAdmin/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

	routes.Setup(app)

    log.Fatal(app.Listen(":8000"))
}

/*
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
*/