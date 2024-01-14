package controllers

import (
	"fmt"
	"github.com/Tibirlayn/GoAdmin/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/Tibirlayn/GoAdmin/pkg/models/account"
)

func GetUser(c *fiber.Ctx) error {
	// создание структуры для записи данных из таблицы TblUser
	var users []account.TblUser
	// подключение к бд и получение всех данных из таблицы TblUser
	result := config.DB.Find(&users)
	if result.Error != nil {
		// Обработать ошибку, если результат неудачен
		fmt.Println("Обработать ошибку, если результат неудачен GetUser")
	}
	// Вернуть данные в формате JSON из таблицы TblUser
	return c.JSON(users)
}

func GetMember(c *fiber.Ctx) error {
	// создание структуры для записи данных из таблицы Member
	var member []account.Member
	// подключение к бд и получение всех данных из таблицы Member
	result := config.DB.Find(&member)
	if result.Error != nil {
		// Обработать ошибку, если результат неудачен
		fmt.Println("Обработать ошибку, если результат неудачен GetMember")
	}
	// Вернуть данные в формате JSON из таблицы Member
	return c.JSON(member)
}

func GetUserBlock(c *fiber.Ctx) error {
	// создание структуры для записи данных из таблицы TblUserBlock
	var userBlock []account.UserBlock
	// подключение к бд и получение всех данных из таблицы TblUserBlock
	result := config.DB.Find(&userBlock)
	if result.Error != nil {
		// Обработать ошибку, если результат неудачен
		fmt.Println("Обработать ошибку, если результат неудачен GetUserBlock")
	}
	// Вернуть данные в формате JSON из таблицы TblUserBlock
	return c.JSON(userBlock)
}

func GetUserBlack(c *fiber.Ctx) error {
	// создание структуры для записи данных из таблицы TblUserBlack
	var userBlack []account.UserBlack
	// подключение к бд и получение всех данных из таблицы TblUserBlack
	result := config.DB.Find(&userBlack)
	if result.Error != nil {
		// Обработать ошибку, если результат неудачен
		fmt.Println("Обработать ошибку, если результат неудачен GetUserBlack")
	}
	// Вернуть данные в формате JSON из таблицы TblUserBlack
	return c.JSON(userBlack)
}

func GetUserAdmin(c *fiber.Ctx) error {
	// создание структуры для записи данных из таблицы TblUserAdmin
	var userAdmin []account.UserAdmin
	// подключение к бд и получение всех данных из таблицы TblUserAdmin
	result := config.DB.Find(&userAdmin)
	if result.Error != nil {
		// Обработать ошибку, если результат неудачен
		fmt.Println("Обработать ошибку, если результат неудачен GetUserAdmin")
	}
	// Вернуть данные в формате JSON из таблицы TblUserAdmin
	return c.JSON(userAdmin)
}

func GetSearchUser(c *fiber.Ctx) error {
	value := c.Params("value")

	
	fmt.Println(value)
	return nil
}