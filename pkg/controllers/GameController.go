package controllers

import (
	"fmt"

	"github.com/Tibirlayn/GoAdmin/pkg/config"
	"github.com/Tibirlayn/GoAdmin/pkg/models/game"
	"github.com/gofiber/fiber/v2"
)

func GetPc(c *fiber.Ctx) error {
	if DB, err := config.GameConfiguration(); err != nil {
		return err
	} else {
		var pc []game.Pc
		if result := DB.Find(&pc); result.Error != nil {
			fmt.Println(err)
		} 
			return c.JSON(pc)
	}
}