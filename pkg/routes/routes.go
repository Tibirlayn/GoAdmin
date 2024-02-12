package routes

import (
	"strconv"

	"github.com/Tibirlayn/GoAdmin/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register) // регистрация
	app.Post("/api/login", controllers.Login)       // логин
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
	app.Get("/api/pc-info", func(c *fiber.Ctx) error { // получить информацию о персонаже
		namePc := c.Query("mNm")
		if namePc == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Name is empaty",
			})
		}
		return controllers.GetPcInfo(c, namePc)
	})
	app.Get("/api/top-pc", controllers.GetTopPcByLevel)     // Запрос на просмотр ТОП 100 игроков по уровню:
	app.Get("/api/top-pc-gold", controllers.GetTopPcbyGold) // Запрос на просмотр ТОП 100 игроков по количеству золота:

	// ParmController
	app.Get("/api/drop-boss", controllers.GetInfoBossDrop)              // Просмотр всех предметов у монстра
	app.Get("/api/specific-proc-item", controllers.GetSpecificProcItem) // Просмотр координат печатей телепорта
	app.Get("/api/gold-chest", controllers.GetGoldChest)                // Посмотреть дроп из золотого/изумрудного сундука + шансы

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
	app.Get("/api/refine-by-name", func(c *fiber.Ctx) error {
		pageNamber, err := strconv.Atoi(c.Query("page"))
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

		return controllers.GetRefineByName(c, pageNamber, limitCnt)
	})
	app.Get("/api/item-resource", func(c *fiber.Ctx) error {
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

		return controllers.GetItemResource(c, pageNumber, limitCnt)
	})
	app.Get("/api/monster-resource", func(c *fiber.Ctx) error {
		pageNumber, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			return err
		}

		limitCnt, err := strconv.Atoi(c.Query("limitCnt"))
		if err != nil {
			return err
		}

		return controllers.GetMonsterResource(c, pageNumber, limitCnt)
	})
	app.Get("/api/top-battle", func(c *fiber.Ctx) error {
		pageNumber, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			return err
		}

		limitCnt, err := strconv.Atoi(c.Query("limitCnt"))
		if err != nil {
			return err
		}

		return controllers.GetTopBattle(c, pageNumber, limitCnt)
	})

	// Просмотр дропа из сундуков
	app.Get("/api/drop-chests", func(c *fiber.Ctx) error {
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

		return controllers.GetDropFromChests(c, pageNumber, limitCnt)
	})

	// Добавить Крафт
	app.Post("/api/add-craft", controllers.PostAddCraft)

	// BillingController
	app.Post("/api/add-gift", controllers.PostGift)               // добавить 1 подарок на аккаунт
	app.Post("/api/add-gift-all", controllers.PostGiftAll)        // добавить всем персонажам подарок
	app.Post("/api/add-gift-pc", controllers.PostGiftPcName)      // добавить подаро по имени персонажа
	app.Delete("/api/delete-all-gift", controllers.DeleteAllGift) // удалить все подарки
}
