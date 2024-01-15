package routes

import (
	"github.com/Tibirlayn/GoAdmin/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register) // регистрация 
	app.Post("/api/login", controllers.Login) // логин
	app.Get("/api/user", controllers.User)	
	app.Post("/api/logout", controllers.Logout) // выход

	// AccountController
	app.Get("/api/get-user", controllers.GetUser)
	app.Get("/api/member", controllers.GetMember)
	app.Get("/api/user-block", controllers.GetUserBlock)
	app.Get("/api/user-black", controllers.GetUserBlack)
	app.Get("/api/admin", controllers.GetUserAdmin)
	app.Get("/api/search-user/:value", controllers.GetSearchUser) // поиск по id / namePc / nameLogin / email 

	// GameController
	app.Get("/api/pc", controllers.GetPc)
	app.Get("/api/user-pc/:idUser", controllers.GetUserPc)
	app.Get("/api/pc-info", controllers.GetPcInfo)



	// BillingController
	// ParmController
}