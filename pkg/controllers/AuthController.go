package controllers

import (
	"github.com/Tibirlayn/GoAdmin/pkg/config"
	"github.com/Tibirlayn/GoAdmin/pkg/models/account"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	connection_DB, err := config.AccountConfiguration()
    if err != nil {
        return err
    }

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := account.User {
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	if err := connection_DB.Create(&user).Error; err != nil {
		return err
	}

	// Получаем объект базы данных gorm.DB и отложенно закрываем его соединение
	sqlDB, err := connection_DB.DB()
    if err != nil {
        // Обработка ошибки
    }

    if err := sqlDB.Close(); err != nil {
        // Обработка ошибки закрытия соединения
    }

	return c.JSON(user)
}