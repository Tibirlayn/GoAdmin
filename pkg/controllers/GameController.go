package controllers

import (
	"fmt"

	"github.com/Tibirlayn/GoAdmin/pkg/config"
	"github.com/Tibirlayn/GoAdmin/pkg/models/game"
	"github.com/gofiber/fiber/v2"
)

type PcInfo struct {
	Pc []game.Pc
	PcState []game.PcState
	PcInventory []game.PcInventory
	PcStore []game.PcStore
}

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

func GetUserPc(c *fiber.Ctx) error {
	var idUser = c.Params("idUser")
	var pc []game.Pc
	if DB, err := config.GameConfiguration(); err != nil {
		return err
	} else {
		if result := DB.Where("mOwner = ?", idUser).Find(&pc); result.Error != nil {
			// Обработать ошибку, если результат неудачен
			fmt.Println("Обработать ошибку, если результат неудачен getUserPc") 
		}
		return c.JSON(pc)
	}
}

func GetPcInfo(c *fiber.Ctx) error {
	var name = c.Query("mNm") // имя персонажа
	var pc PcInfo
	if DB, err := config.GameConfiguration(); err != nil {
		return err
	} else {
		if resultPc := DB.Where("mNm = ?", name).Find(&pc.Pc); resultPc.Error != nil {
			fmt.Println("Обработать ошибку, если результат неудачен GetPcInfo > if > resultPc") 
		} 

		idPc := pc.Pc[0].MNo
		idUser := pc.Pc[0].MOwner

		if resultPcState := DB.Where("mNo = ?", idPc).Find(&pc.PcState); resultPcState.Error != nil {
			fmt.Println("Обработать ошибку, если результат неудачен GetPcInfo > if > resultPcState") 
		} 
		if resultPcInventory := DB.Where("mPcNo = ?", idPc).Find(&pc.PcInventory); resultPcInventory.Error != nil {
			fmt.Println("Обработать ошибку, если результат неудачен GetPcInfo > if > resultPcInventory") 
		}
		if resultPcStore := DB.Where("mUserNo = ?", idUser).Find(&pc.PcStore); resultPcStore.Error != nil {
			fmt.Println("Обработать ошибку, если результат неудачен GetPcInfo > if > resultPcStore") 
		}

	}
	return c.JSON(pc)
}
