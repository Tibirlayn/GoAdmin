package controllers

import (
	"fmt"
	"time"
	"github.com/Tibirlayn/GoAdmin/pkg/config"
	"github.com/Tibirlayn/GoAdmin/pkg/models/account"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

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

	user := account.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	if err := connection_DB.Create(&user).Error; err != nil {
		return err
	}

	// Получаем объект базы данных gorm.DB и отложенно закрываем его соединение
	sqlDB, err := connection_DB.DB()
	if err != nil {
		// Обработка ошибки
		panic(err)
	}

	if err := sqlDB.Close(); err != nil {
		// Обработка ошибки закрытия соединения
		panic(err)
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user account.User

	connection_DB, err := config.AccountConfiguration()
	if err != nil {
		return err
	}

	connection_DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	// Получаем объект базы данных gorm.DB и отложенно закрываем его соединение
	sqlDB, err := connection_DB.DB()
	if err != nil {
		// Обработка ошибки
		panic(err)
	}

	if err := sqlDB.Close(); err != nil {
		// Обработка ошибки закрытия соединения
		panic(err)
	}

	// Генерация токена JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Устанавливаем срок действия токена на 72 часа

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
	//return c.JSON(user)
}

func User(c *fiber.Ctx) error {
	// Получаем токен из cookie
	cookie := c.Cookies("jwt")

	// Пытаемся распарсить токен с помощью секретного ключа
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что тип подписи подходит для нашего секретного ключа
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		// Возвращаем секретный ключ
		fmt.Println(2)
		return []byte(SecretKey), nil
	})
	// Проверяем ошибки при парсинге токена
	if err != nil {
		return err // Обработка ошибки
	}
	
	// Проверяем валидность токена
	if token.Valid {
		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok {
			return fmt.Errorf("invalid claims")
		}
		// Распарсиваем ID пользователя из токена
		userID, ok := (*claims)["id"]
		if !ok {
			return fmt.Errorf("id claim not found or invalid")
		}
		var user account.User
		// Получаем информацию о пользователе из базы данных
		connection_DB, err := config.AccountConfiguration()
		if err != nil {
			return err
		}

		connection_DB.Where("id = ?", userID).First(&user)

		return c.JSON(user)
	} else {
		return fmt.Errorf("invalid token")
	}
}

func Logout(c *fiber.Ctx) error{
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
