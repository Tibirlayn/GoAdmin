package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Tibirlayn/GoAdmin/pkg/config"
	"github.com/Tibirlayn/GoAdmin/pkg/models/billing"
	"github.com/Tibirlayn/GoAdmin/pkg/models/game"
	"github.com/Tibirlayn/GoAdmin/pkg/models/parm"
	"github.com/gofiber/fiber/v2"
)

// Удалить падарки всем персонажам
func DeleteAllGift(c *fiber.Ctx) error {
	BillingDB, err := config.BillingConfiguration()
	if err != nil {
		return err
	}

	tx := BillingDB.Begin()
	if err := BillingDB.Unscoped().Where("1 = 1").Delete(&billing.SysOrderList{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "data delete",
	})
}

// Выдать только один подарок на сервер
func PostGift(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, _ := strconv.ParseInt(data["id"], 10, 64)                        /* — ID сообщения от администратора */
	svr, _ := strconv.ParseInt(data["svr"], 10, 16)                      /* — Номер вашего сервера */
	itemid, _ := strconv.Atoi(data["itemid"])                            /* — Номер предмета (подарка) */
	cnt, _ := strconv.Atoi(data["cnt"])                                  /* — Количество */
	aperiod, _ := strconv.Atoi(data["aperiod"])                          /* — Доступный период (сколько будет лежать в подароках) */
	pperiod, _ := strconv.Atoi(data["pperiod"])                          /* — Практический период (количество времени которое будет у предмета после получения)*/
	binding, _ := strconv.ParseUint(data["binding"], 10, 8)              /* — Под замком предмет или нет (Нет = 0 | Да = 1) */
	limitedDate, _ := time.Parse("2006-01-02 15:04:05", data["limited"]) /* — Ограниченная дата */
	status, _ := strconv.ParseUint(data["status"], 10, 8)                /* — Статус предмета */

	BillingDB, err := config.BillingConfiguration()
	if err != nil {
		return err
	}
	GameDB, err := config.GameConfiguration()
	if err != nil {
		return err
	}

	var owners []int

	if err := GameDB.Model(&game.Pc{}).Distinct().Pluck("mOwner", &owners).Error; err != nil {
		return err
	}

	tx := BillingDB.Begin()
	for _, owner := range owners {
		giftPc := billing.SysOrderList{
			MSysID:           id,
			MUserNo:          owner,
			MSvrNo:           int16(svr),
			MItemID:          itemid,
			MCnt:             cnt,
			MAvailablePeriod: aperiod,
			MPracticalPeriod: pperiod,
			MBindingType:     uint8(binding),
			MLimitedDate:     limitedDate,
			MItemStatus:      uint8(status),
		}
		if err := BillingDB.
			Omit("mRegDate", "mReceiptDate", "mReceiptPcNo", "mRecepitPcNm").
			Create(&giftPc).Error; err != nil {
			tx.Rollback() // Откатить транзакцию при возникновении ошибки
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Gifts added",
	})

}

// выдать подарок всем персонажам
func PostGiftAll(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, _ := strconv.ParseInt(data["id"], 10, 64)                        /* — ID сообщения от администратора */
	svr, _ := strconv.ParseInt(data["svr"], 10, 16)                      /* — Номер вашего сервера */
	itemid, _ := strconv.Atoi(data["itemid"])                            /* — Номер предмета (подарка) */
	cnt, _ := strconv.Atoi(data["cnt"])                                  /* — Количество */
	aperiod, _ := strconv.Atoi(data["aperiod"])                          /* — Доступный период (сколько будет лежать в подароках) */
	pperiod, _ := strconv.Atoi(data["pperiod"])                          /* — Практический период (количество времени которое будет у предмета после получения)*/
	binding, _ := strconv.ParseUint(data["binding"], 10, 8)              /* — Под замком предмет или нет (Нет = 0 | Да = 1) */
	limitedDate, _ := time.Parse("2006-01-02 15:04:05", data["limited"]) /* — Ограниченная дата */
	status, _ := strconv.ParseUint(data["status"], 10, 8)                /* — Статус предмета */

	BillingDB, err := config.BillingConfiguration()
	if err != nil {
		return err
	}
	GameDB, err := config.GameConfiguration()
	if err != nil {
		return err
	}

	var owners []int

	if err := GameDB.Model(&game.Pc{}).Where("mDelDate IS NULL").Pluck("mOwner", &owners).Error; err != nil {
		fmt.Println("Error GameDB.Model(&game.Pc{})")
		return err
	}

	tx := BillingDB.Begin()
	for _, owner := range owners {
		giftPc := billing.SysOrderList{
			MSysID:           id,
			MUserNo:          owner,
			MSvrNo:           int16(svr),
			MItemID:          itemid,
			MCnt:             cnt,
			MAvailablePeriod: aperiod,
			MPracticalPeriod: pperiod,
			MBindingType:     uint8(binding),
			MLimitedDate:     limitedDate,
			MItemStatus:      uint8(status),
		}
		if err := BillingDB.Omit("mRegDate", "mReceiptDate", "mReceiptPcNo", "mRecepitPcNm").Create(&giftPc).Error; err != nil {
			tx.Rollback() // Откатить транзакцию при возникновении ошибки
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Gifts added",
	})
}

// выдать подарок персонажу
func PostGiftPcName(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, _ := strconv.ParseInt(data["id"], 10, 64)                        /* — ID сообщения от администратора */
	svr, _ := strconv.ParseInt(data["svr"], 10, 16)                      /* — Номер вашего сервера */
	itemid, _ := strconv.Atoi(data["itemid"])                            /* — Номер предмета (подарка) */
	mNm, _ := data["mNm"]                                                /* — Имя персонажа */
	cnt, _ := strconv.Atoi(data["cnt"])                                  /* — Количество */
	aperiod, _ := strconv.Atoi(data["aperiod"])                          /* — Доступный период (сколько будет лежать в подароках) */
	pperiod, _ := strconv.Atoi(data["pperiod"])                          /* — Практический период (количество времени которое будет у предмета после получения)*/
	binding, _ := strconv.ParseUint(data["binding"], 10, 8)              /* — Под замком предмет или нет (Нет = 0 | Да = 1) */
	limitedDate, _ := time.Parse("2006-01-02 15:04:05", data["limited"]) /* — Ограниченная дата */
	status, _ := strconv.ParseUint(data["status"], 10, 8)                /* — Статус предмета */

	// Подключение к БД Billing
	BillingDB, err := config.BillingConfiguration()
	if err != nil {
		return err
	}
	// Подключение к БД GameDB
	GameDB, err := config.GameConfiguration()
	if err != nil {
		return err
	}
	// Создаем массив
	var owners []int

	// поиск персонажа в БД GameDB и записываем в массив owners
	// (Pluck - этот метод производит выборку определенного столбца ("mOwner") из результата запроса и сохраняет значения в срез "owners")
	if err := GameDB.Model(&game.Pc{}).Where("mNm = ?", mNm).Pluck("mOwner", &owners).Error; err != nil {
		fmt.Println("Error GameDB.Model(&game.Pc{})")
		return err
	}

	tx := BillingDB.Begin() // создаем транзакцию
	for _, owner := range owners {
		giftPc := billing.SysOrderList{
			MSysID:           id,
			MUserNo:          owner,
			MSvrNo:           int16(svr),
			MItemID:          itemid,
			MCnt:             cnt,
			MAvailablePeriod: aperiod,
			MPracticalPeriod: pperiod,
			MBindingType:     uint8(binding),
			MLimitedDate:     limitedDate,
			MItemStatus:      uint8(status),
		}
		// omit игнорирует на запись (данные поля заполняются автоматически в бд или их не нужно заполнять)
		if err := BillingDB.Omit("mRegDate", "mReceiptDate", "mReceiptPcNo", "mRecepitPcNm").Create(&giftPc).Error; err != nil {
			tx.Rollback() // Откатить транзакцию при возникновении ошибки
			return err
		}
	}

	// сохраняем изменения
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Gifts added",
	})
}

// SQL Запрос. Добавить предмет в ШОП -> нужно тест
func PostAddShopItem(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	itemID, _ := strconv.Atoi(data["itemID"])                // DECLARE @ItemID INT = 8531 /* ID предмета*/
	count, _ := strconv.Atoi(data["count"])                  // DECLARE @ICount INT = 1 /* Количество предметов */
	name, _ := data["name"]                                  // DECLARE @IName VARCHAR(40) = 'Особое Зельe Жизни' /* Название предмета */
	desc, _ := data["desc"]                                  // DECLARE @IDesc VARCHAR(500) = 'Средство, восстанавливающее большое количество здоровья.' /* Описание предмета */
	price, _ := strconv.Atoi(data["price"])                  // DECLARE @IPrice INT = 300 /* Цена */
	status, _ := data["status"]                              // DECLARE @Istatus INT = 1 /* Статус 0 проклятый, 1-обычный, 2-благой */
	cat, _ := strconv.ParseInt(data["itemCategory"], 10, 16) // DECLARE @ICat INT = 3 /* Вкладка шопа 1,2,3,4 */
	day, _ := strconv.Atoi("AvailablePeriod")                // DECLARE @IDay INT = 30 /* Время на предмете в днях*/
	hour, _ := strconv.Atoi("PracticalPeriod")               // DECLARE @IHour INT = 0 /* Время эффекта предмета в днях */
	svrNo, _ := strconv.ParseInt(data["svr"], 10, 16)        // DECLARE @SvrNo INT = 1164 /* Номер сервера */

	BillingDB, err := config.BillingConfiguration()
	if err != nil {
		return err
	}

	ParmDB, err := config.ParmConfiguration()
	if err != nil {
		return err
	}

	var item parm.Item

	// Проверка, существует ли запись с указанным ItemID и IIsCharge = 1
	result := ParmDB.First(&item, "IID = ? AND IIsCharge = ?", itemID, 1)

	if result.RowsAffected == 0 {
		// Если записи не существует, выполняем обновление
		ParmDB.Model(&parm.Item{}).Where("IID = ?", itemID).Update("IIsCharge", 1)
	}

	currentTime := time.Now() /* Текущая дата */
	var maxGoldenID int64
	var maxOrder int16
	var packageGold = "0"
	var admin = "GoAdmin"
	if err := BillingDB.Model(&billing.GoldItem{}).
		Select("MAX(GoldItemID)").
		Find(&billing.GoldItem{}).
		Scan(&maxGoldenID).Error; err != nil {
		return err
	}
	if err := BillingDB.Model(&billing.CategoryAssign{}).
		Select("MAX(OrderNO)").
		Find(&billing.CategoryAssign{}).
		Scan(&maxOrder).Error; err != nil {
		return err
	}

	maxGoldenID += 1 // получаем max число id в таблице GoldItem и + 1
	maxOrder += 1    // получаем max число id в таблице CategoryAssign и + 1

	if maxGoldenID == 0 {
		maxGoldenID = 1
	}
	if maxOrder == 0 {
		maxOrder = 1
	}

	newTBLGoldItem := billing.GoldItem{
		GoldItemID:        maxGoldenID,
		IID:               itemID,
		ItemName:          name,
		ItemDesc:          desc,
		OriginalGoldPrice: price,
		GoldPrice:         price,
		ItemCategory:      int16(cat),
		IsPackage:         packageGold,
		Status:            status,
		AvailablePeriod:   day,
		Count:             count,
		PracticalPeriod:   hour,
		RegistAdmin:       admin,
	}

	tx := BillingDB.Begin()
	if err := BillingDB.
		Omit("ItemImage", "RegistDate", "RegistIP", "UpdateDate", "UpdateAdmin", "UpdateIP", "ItemNameRUS", "ItemDescRUS").
		Create(&newTBLGoldItem).Error; err != nil {
		tx.Rollback()
		return err
	}

	newTBLCategoryAssign := billing.CategoryAssign{
		CategoryID:  int16(cat),
		GoldItemID:  maxGoldenID,
		Status:      "1",
		OrderNO:     maxOrder,
		RegistDate:  currentTime,
		RegistAdmin: admin,
	}

	if err := BillingDB.
		Omit("RegistIP", "UpdateDate", "UpdateAdmin", "UpdateIP").
		Create(&newTBLCategoryAssign).Error; err != nil {
		tx.Rollback()
		return err
	}

	newTBLGoldItemSupportSvr := billing.GoldItemSupportSvr{
		GoldItemID: maxGoldenID,
		MSvrNo:     int16(svrNo),
	}

	if err := BillingDB.Create(&newTBLGoldItemSupportSvr).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "добавлен новый шоп предмет",
	})
}

// DECLARE @Date DATETIME SET @Date = GETDATE() /* Сегодняшняя дата */
// DECLARE @GIid INT = (SELECT MAX(GoldItemID) FROM TBLGoldItem) + 1
// DECLARE @GIOrder INT = (SELECT MAX(OrderNO) FROM TBLCategoryAssign) + 1
// IF @GIid IS NULL SET @GIid = 1
// IF @GIOrder IS NULL SET @GIOrder = 1
// INSERT INTO TBLGoldItem (GoldItemID, IID, ItemName, ItemDesc, OriginalGoldPrice, GoldPrice, ItemCategory, IsPackage, Status, AvailablePeriod, Count, PracticalPeriod, RegistAdmin)
// VALUES (@GIid, @ItemID, @IName, @IDesc, @IPrice, @IPrice, @ICat, 0, @Istatus, @IDay, @ICount, @IHour, 'R2Genius')

// INSERT INTO TBLCategoryAssign (CategoryID, GoldItemID, Status, OrderNO, RegistDate, RegistAdmin)
// VALUES (@ICat, @GIid, 1, @GIOrder, @Date, 'R2Genius')

// INSERT INTO TBLGoldItemSupportSvr(GoldItemID, mSvrNo)
// VALUES (@GIid, @SvrNo)
