package routes

import (
	"strconv"

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
	app.Get("/api/search-user", controllers.GetSearchUser) // поиск по id / namePc / nameLogin / email 

	// GameController
	app.Get("/api/pc", controllers.GetPc) // получить персонажа
	app.Get("/api/user-pc/:idUser", controllers.GetUserPc)
	app.Get("/api/pc-info", controllers.GetPcInfo)

	
	// ParmController
	app.Get("/api/drop-boss", controllers.GetInfoBossDrop) // Просмотр всех предметов у монстра
	app.Get("/api/specific-proc-item", controllers.GetSpecificProcItem) // Просмотр координат печатей телепорта 
	//app.Get("/api/refine", controllers.GetRefine) // Просмотр DT_Refine
	// Просмотр DT_Refine
	app.Get("/api/refine", func(c *fiber.Ctx) error {
		pageNumber, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid page number",
			})
		}
		limitCnt, err := strconv.Atoi(c.Query("limitCnt"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid limitCnt",
			})
		}

		return controllers.GetRefine(c, pageNumber, limitCnt)
	}) 
	
	// BillingController
	app.Post("/api/add-gift", controllers.PostGift) // добавить 1 подарок на аккаунт
	app.Post("/api/add-gift-all", controllers.PostGiftAll) // добавить всем персонажам подарок
	app.Post("/api/add-gift-pc", controllers.PostGiftPcName) // добавить подаро по имени персонажа
	app.Delete("/api/delete-all-gift", controllers.DeleteAllGift) // удалить все подарки 
}