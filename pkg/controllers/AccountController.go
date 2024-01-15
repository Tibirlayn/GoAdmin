package controllers

import (
	"fmt"
	"github.com/Tibirlayn/GoAdmin/pkg/config"
	"github.com/Tibirlayn/GoAdmin/pkg/models/account"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type SearchResult struct {
	Member []account.Member
	User []account.TblUser
}

func GetUser(c *fiber.Ctx) error {
	if DB, err := config.AccountConfiguration(); err != nil {
		return err
	} else {
		// создание структуры для записи данных из таблицы TblUser
		var users []account.TblUser
		// подключение к бд и получение всех данных из таблицы TblUser
		if result := DB.Find(&users); result.Error != nil {
			// Обработать ошибку, если результат неудачен
			fmt.Println("Обработать ошибку, если результат неудачен GetUser")
		}
		// Вернуть данные в формате JSON из таблицы TblUser
		return c.JSON(users)
	}
}

func GetMember(c *fiber.Ctx) error {
	if DB, err := config.AccountConfiguration(); err != nil {
		return err
	} else {
		// создание структуры для записи данных из таблицы Member
		var member []account.Member
		// подключение к бд и получение всех данных из таблицы Member
		if result := DB.Find(&member); result.Error != nil {
			// Обработать ошибку, если результат неудачен
			fmt.Println("Обработать ошибку, если результат неудачен GetMember")
		}
		// Вернуть данные в формате JSON из таблицы Member
		return c.JSON(member)
	}
}

func GetUserBlock(c *fiber.Ctx) error {
	if DB, err := config.AccountConfiguration(); err != nil {
		return err
	} else {
		// создание структуры для записи данных из таблицы TblUserBlock
		var userBlock []account.UserBlock
		// подключение к бд и получение всех данных из таблицы TblUserBlock
		if result := DB.Find(&userBlock); result.Error != nil {
			// Обработать ошибку, если результат неудачен
			fmt.Println("Обработать ошибку, если результат неудачен GetUserBlock")
		}
		// Вернуть данные в формате JSON из таблицы TblUserBlock
		return c.JSON(userBlock)
	}
}

func GetUserBlack(c *fiber.Ctx) error {
	if DB, err := config.AccountConfiguration(); err != nil {
		return err
	} else {
		var userBlack []account.UserBlack
		if result := DB.Find(&userBlack); result.Error != nil {
			// Обработать ошибку, если результат неудачен
			fmt.Println("Обработать ошибку, если результат неудачен GetUserBlack")
		}
		// Вернуть данные в формате JSON из таблицы TblUserBlack
		return c.JSON(userBlack)
	}
}

func GetUserAdmin(c *fiber.Ctx) error {
	if DB, err := config.AccountConfiguration(); err != nil {
		return err
	} else {
		// создание структуры для записи данных из таблицы TblUserAdmin
		var userAdmin []account.UserAdmin
		// подключение к бд и получение всех данных из таблицы TblUserAdmin
		if result := DB.Find(&userAdmin); result.Error != nil {
			// Обработать ошибку, если результат неудачен
			fmt.Println("Обработать ошибку, если результат неудачен GetUserAdmin")
		}
		// Вернуть данные в формате JSON из таблицы TblUserAdmin
		return c.JSON(userAdmin)
	}
}
// поиск по id / namePc / nameLogin / email 
func GetSearchUser(c *fiber.Ctx) error {
	if DB, err := config.AccountConfiguration(); err != nil {
		return err
	} else {
		value := c.Params("value")
		var result SearchResult
		//memberUserId := DB.Table(member.TableName()).Where("MUserId = ?", value).First(&member)
		memberResult := DB.Where("mUserId = ? OR email = ?", value, value).Find(&result.Member)
		if num, err := strconv.Atoi(value); err != nil {
			DB.Where("mUserId = ?", value).Find(&result.User)
			fmt.Println("Ошибка преобразования строки в число:", err)
		} else {
			DB.Where("mUserNo = ? OR mUserId = ?", num, value).Find(&result.User)
		}
		
		if memberResult.Error != nil {
			// Обработать ошибку, если результат неудачен
			fmt.Println("Ошибка при выполнении запроса к базе данных:", memberResult.Error)
			return c.Status(500).SendString("Ошибка при выполнении запроса к базе данных")
		}
		return c.JSON(result)
	}
}

/*
func GetUserPc(c fiber.Ctx) error {
	if DB, err := config.AccountConfiguration(); err != nil {
		return err
	} else {
		value := c.Params("value") 
		var pc := 
	}
	return nil
}
*/
