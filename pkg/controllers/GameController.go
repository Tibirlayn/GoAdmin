package controllers

import (
	"errors"
	"fmt"

	"github.com/Tibirlayn/GoAdmin/pkg/config"
	"github.com/Tibirlayn/GoAdmin/pkg/models/game"
	"github.com/gofiber/fiber/v2"
)

type PcInfo struct {
	Pc          []game.Pc
	PcState     []game.PcState
	PcInventory []game.PcInventory
	PcStore     []game.PcStore
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
	// имя персонажа
	var name = c.Query("mNm")
	if name == "" {
		return errors.New("Введите значение персонажа")
	}
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

// Запрос на просмотр ТОП 100 игроков по уровню:
func GetTopPcByLevel(c *fiber.Ctx) error {

	GameDB, err := config.GameConfiguration()
	if err != nil {
		return err
	}

	var result []struct {
		ID      int    `gorm:"column:mNo"`      // id персонажа
		Class   int8   `gorm:"column:mClass"`   // класс персонажа
		Name    string `gorm:"column:mNm"`      // имя персонажа
		Level   int16  `gorm:"column:mLevel"`   // уровень персонажа
		Chaotic int    `gorm:"column:mChaotic"` // рейтинг персонажа
		PkCnt   int    `gorm:"column:mPkCnt"`   // кол-во убийств
	}

	if err := GameDB.Table("TblPc AS a").
		Select("TOP 100 a.mNo AS ID, a.mClass AS Class, RTRIM(a.mNm) AS Name, b.mLevel AS Level, b.mChaotic AS Chaotic, b.mPkCnt AS PkCnt").
		Joins("JOIN TblPcState AS b ON a.mNo = b.mNo").
		Where("a.mNo > ? AND LEFT(a.mNm, 1) <> ?", 0, ",").
		Order("b.mLevel DESC").
		Scan(&result).Error; err != nil {
		return err
	}

	return c.JSON(result)
}

// Запрос на просмотр ТОП 100 игроков по количеству золота:
func GetTopPcbyGold(c *fiber.Ctx) error {

	GameDB, err := config.GameConfiguration()
	if err != nil {
		return err
	}

	var result []struct {
		MOwner    int    `gorm:"column:mOwner"` // аккаунт персонажа
		MSerialNo int64  `gorm:"column:mSerialNo"`
		Name      string `gorm:"column:mNm"`     // имя персонажа
		MPcNo     int    `gorm:"column:mPcNo"`   // id персонажа
		MItemNo   int    `gorm:"column:mItemNo"` // id предмета
		MCnt      int    `gorm:"column:mCnt"`    // кол-во
	}

	if err := GameDB.Table("TblPc AS a").
		Select("TOP 100 a.mOwner AS MOwner, b.mSerialNo AS MSerialNo, RTRIM(a.mNm) AS Name, b.mPcNo AS MPcNo, b.mItemNo AS MItemNo, b.mCnt AS MCnt").
		Joins("INNER JOIN TblPcInventory AS b ON b.mPcNo = a.mNo").
		Where("b.mItemNo = ? AND LEFT (a.mNm, 1) <> ?", 409, ",").
		Order("b.mCnt DESC").
		Scan(&result).Error; err != nil {
		return err
	}

	return c.JSON(result)
}
