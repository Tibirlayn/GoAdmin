package controllers

import (
	//	"encoding/json"
	"errors"
	"fmt"

	"github.com/Tibirlayn/GoAdmin/pkg/config"
	"github.com/Tibirlayn/GoAdmin/pkg/models/parm"
	"github.com/gofiber/fiber/v2"
)

type ItemResult struct {
	Boss      string
	NameGroup string
	NameItem  string
	ItemDesc  string
	Event     string
}

type SpecificItem struct {
	ItemIID  int
	ItemName string
	X        int64
	Y        int64
	Z        int64
}

func GetInfoBossDrop(c *fiber.Ctx) error {
	mid := c.Query("MID")
	name := c.Query("MName")
	if mid == "" && name == "" {
		return errors.New("Введите id монстра или название монстра! Поле не может быть пустым!")
	} else if mid != "" && name != "" {
		return errors.New("Введите что-то одно: id монстра или название монстра!")
	}
	if DB, err := config.ParmConfiguration(); err != nil {
		return err
	} else {
		results := []ItemResult{}

		query := `
			SELECT DISTINCT
				mo.MName AS 'Boss',
				tdg.DName AS 'NameGroup',
				it.IName AS 'NameItem',
				tpis.mDesc AS 'ItemDesc',
				CASE
					WHEN evo.mObjID IS NULL THEN 'Not event item'
					ELSE 'Event item'
				END AS 'Event'
				FROM DT_Monster AS mo
				INNER JOIN DT_MonsterDrop AS md ON mo.MID = md.MID
				INNER JOIN DT_DropGroup AS dg ON md.DGroup = dg.DGroup
				INNER JOIN TP_DropGroup AS tdg ON md.DGroup = tdg.DGroup
				INNER JOIN DT_DropItem AS di ON dg.DDrop = di.DDrop
				INNER JOIN TP_ItemStatus AS tpis ON di.DStatus = tpis.mStatus
				INNER JOIN DT_Item AS it ON di.DItem = it.IID
				LEFT JOIN TblEventObj AS evo ON it.IID = evo.mObjID
				WHERE
					mo.MID = ? OR mo.MName = ? AND
					di.DIsEvent = 0
		`
		if err := DB.Raw(query, mid, name).Scan(&results).Error; err != nil {
			// Handle error
			return err
		}

		// проверка для себя
		if false {
			for _, item := range results {
				fmt.Println(item)
			}
		}
		return c.JSON(results)

	}
}

func GetSpecificProcItem(c *fiber.Ctx) error {
	name := c.Query("NameItem")
	fmt.Println(name)

	if DB, err := config.ParmConfiguration(); err != nil {
		return err
	} else {
		result := []SpecificItem{}

		query := `
		SELECT
		a.mIID as 'ItemIID',
		b.IName as 'ItemName',
		a.mAParam as 'X',
		a.mBParam as 'Y',
		a.mCParam as 'Z'
		FROM TblSpecificProcItem as a
		INNER JOIN DT_Item as b on (a.mIID = b.IID)
		WHERE b.IName LIKE ?
		`

		if err := DB.Raw(query, "%"+name+"%").Scan(&result).Error; err != nil {
			return err
		}

		return c.JSON(result)
	}
}

// Просмотр DT_Refine > сделать поиск по названию
func GetRefine(c *fiber.Ctx, pageNumber int, limitCnt int) error {

	ParmDB, err := config.ParmConfiguration()
	if err != nil {
		return err
	}

	// Рассчитываем смещение (offset) и лимит записей на основе номера страницы
	limit := limitCnt
	offset := (pageNumber - 1) * limit

	var results []struct {
		ID               int     `gorm:"column:RID"`
		ReceivedItem     int     `gorm:"column:RItemID0"`
		ReceivedItemName string  `gorm:"column:IName"`
		RecipeItemID     int     `gorm:"column:RItemID"`
		RecipeItemName   string  `gorm:"column:IName"`
		SuccessChance    float64 `gorm:"column:RSuccess"`
	}

	if err := ParmDB.Table("DT_Refine a").
		Select("a.RID as ID, a.RItemID0, b.IName as ReceivedItemName, c.RItemID as RecipeItemID, b1.IName as RecipeItemName, a.RSuccess as SuccessChance").
		Joins("INNER JOIN DT_Item as b ON a.RItemID0 = b.IID").
		Joins("INNER JOIN DT_RefineMaterial as c ON a.RID = c.RID").
		Joins("INNER JOIN DT_Item as b1 ON c.RItemID = b1.IID").
		Offset(offset).
		Limit(limit).
		Scan(&results).Error; err != nil {
		return err
	}

	return c.JSON(results)
}

// Названия предметов из рецепта в одну строку
func GetRefineByName(c *fiber.Ctx, pageNumber int, limitCnt int) error {
	ParmBD, err := config.ParmConfiguration()
	if err != nil {
		return err
	}

	limit := limitCnt
	offset := (pageNumber - 1) * limit

	var results []struct {
		ID               int     `gorm:"column:RID"`
		ReceivedItem     int     `gorm:"column:RItemID0"`
		ReceivedItemName string  `gorm:"column:IName"`
		RecipeItemName   string  `gorm:"column:IName"`
		SuccessChance    float64 `gorm:"column:RSuccess"`
	}

	if err := ParmBD.Table("DT_Refine a").
		Select("a.RID, a.RItemID0 as ReceivedItem, b.IName as ReceivedItemName, STRING_AGG(b1.IName, ', ') as RecipeItemName, a.RSuccess as SuccessChance").
		Joins("INNER JOIN DT_Item b ON a.RItemID0 = b.IID").
		Joins("INNER JOIN DT_RefineMaterial c ON a.RID = c.RID").
		Joins("INNER JOIN DT_Item b1 ON c.RItemID = b1.IID").
		Group("a.RID, a.RItemID0, b.IName, a.RSuccess").
		Order("RID").
		Offset(offset).
		Limit(limit).
		Scan(&results).Error; err != nil {
		return err
	}

	return c.JSON(results)
}

// Запрос на просмотр DT_ItemResource:
func GetItemResource(c *fiber.Ctx, pageNumber int, limitCnt int) error {
	ParmDB, err := config.ParmConfiguration()
	if err != nil {
		return err
	}

	limit := limitCnt
	offset := (pageNumber - 1) * limit

	var results []struct {
		RID       int    `gorm:"column:RID"`
		RType     int    `gorm:"column:Type"`
		ROwnerID  int    `gorm:"column:OwnerID"`
		IName     string `gorm:"column:Name"`
		RFileName string `gorm:"column:FileName"`
		RPosX     int    `gorm:"column:CoordinateX"`
		RPosY     int    `gorm:"column:CoordinateY"`
	}

	if err := ParmDB.Table("DT_ItemResource a").
		Select("a.RID as RID, a.RType as Type, a.ROwnerID as OwnerID, b.IName as Name, " +
			"CASE a.RType WHEN 2 THEN '.DDS:' WHEN 0 THEN 'MODEL:' END as FileType, " +
			"a.RFileName as FileName, a.RPosX as CoordinateX, a.RPosY as CoordinateY").
		Joins("LEFT OUTER JOIN DT_Item b ON b.IID = a.ROwnerID").
		Limit(limit).
		Offset(offset).
		Scan(&results).Error; err != nil {
		return err
	}

	return c.JSON(results)
}

// Запрос на просмотр DT_MonsterResource
func GetMonsterResource(c *fiber.Ctx, pageNumber int, limitCnt int) error {
	ParmDB, err := config.ParmConfiguration()
	if err != nil {
		return err
	}

	limit := limitCnt
	offset := (pageNumber - 1) * limit

	var result []struct {
		RID       int    `gorm:"column:ID"`
		ROwnerID  int    `gorm:"column:OwnerID"`
		MName     string `gorm:"column:Name"`
		RType     string `gorm:"column:NameTexture"`
		RFileName string `gorm:"column:NumberTXT"`
	}

	//CASE DT_MonsterResource.RType WHEN 0 THEN 'TEXTURE:' END AS 'Имя текстуры',
	if err := ParmDB.Table("DT_MonsterResource as a").
		Select("a.RID as ID, a.ROwnerID as OwnerID, b.MName as Name," +
			"CASE a.RType WHEN 0 THEN 'TEXTURE:' END AS NameTexture, a.RFileName as NumberTXT").
		Joins("LEFT JOIN DT_Monster AS b ON b.MID = a.ROwnerID").
		Limit(limit).
		Offset(offset).
		Scan(&result).Error; err != nil {
		return err
	}

	return c.JSON(result)
}

// SQL Запрос. Посмотреть дроп из золотого/изумрудного сундука + шансы -- (в планах)

// SQL Запросы. Просмотр ТОП по БК/БГ и уровню игроков
func GetTopBattle(c *fiber.Ctx, pageNumber int, limitCnt int) error {

	ParmDB, err := config.ParmConfiguration()
	if err != nil {
		return err
	}

	limit := limitCnt
	offset := (pageNumber - 1) * limit

	var result []struct {
		MRanking    int16  `gorm:"column:mRanking"`
		MGuildNm    string `gorm:"column:mGuildNm"`
		MVictoryCnt int    `gorm:"column:mVictoryCnt"`
		MPcNm       string `gorm:"column:mPcNm"`
	}

	if err := ParmDB.Table("TblUnitedGuildWarHerosBattleRanking as a").
		Select("mRanking, mGuildNm, mVictoryCnt, mVictoryCnt, mPcNm").
		Order("mRanking ASC").
		Limit(limit).
		Offset(offset).
		Scan(&result).Error; err != nil {
		return err
	}

	return c.JSON(result)
}

// Запрос на просмотр рейтинга по БГ
func GetRatingGuild(c *fiber.Ctx, pageNumber int, limitCtx int) error {

	/* 	ParmDB, err := config.ParmConfiguration()

		var result []struct {
			mKillCnt
			mGuildNm
			mVictoryCnt
			mGuildPoint
			mSum
		}



	SELECT
	mKillCnt,
	mGuildNm,
	mVictoryCnt,
	mGuildPoint,
	mSum
	FROM
	[FNLParm].[dbo].[temp_ranking_guild]
	ORDER BY
	mVictoryCnt DESC

		return c.JSON(result) */

	return nil
}

// Просмотр дропа из сундуков
func GetDropFromChests(c *fiber.Ctx, pageNumber int, limitCnt int) error {

	limit := limitCnt
	offset := (pageNumber - 1) * limit

	ParmDB, err := config.ParmConfiguration()
	if err != nil {
		return err
	}

	var result []struct {
		MDID        int64   `gorm:"column:MDID"`
		MDRD        int64   `gorm:"column:MDRD"`
		IdBox       int     `gorm:"column:IID"`         // IID Коробки
		NameBox     string  `gorm:"column:IName"`       // Название Коробки
		IdItem      int     `gorm:"column:IID"`         // ID Получаемого предмета
		NameItem    string  `gorm:"column:IName"`       // Название Получаемого предмета
		MDesc       string  `gorm:"column:mDesc"`       // Название группы
		MPerOrRate  float64 `gorm:"column:mPerOrRate"`  // Дроп шанс
		MItemStatus int8    `gorm:"column:mItemStatus"` // Статус предмета
		MCnt        int     `gorm:"column:mCnt"`        // Количество
		MBinding    string  `gorm:"column:mBinding"`    // Предмет Под замком?
		MEffTime    int     `gorm:"column:mEffTime"`    // Время эффекта
		MValTime    int16   `gorm:"column:mValTime"`    // Время предмета
		MMaxResCnt  int     `gorm:"column:mMaxResCnt"`  // MaxResCnt DrawIndex
		MSuccess    float64 `gorm:"column:mSuccess"`    // Шанс DrawIndex
	}

	if err := ParmDB.Table("TblMaterialDrawResult AS a").
		Select("a2.MDID, a.MDRD, a3.IID AS IdBox, b2.IName AS NameBox, a.IID AS IdItem, b.IName AS NameItem, a2.mDesc AS MDesc, a.mPerOrRate AS MPerOrRate, " +
			"a.mItemStatus AS MItemStatus, a.mCnt AS MCnt, CASE a.mBinding WHEN 0 THEN 'YES' WHEN 1 THEN 'NO' END AS MBinding, a.mEffTime AS MEffTime, " +
			"a.mValTime AS MValTime, a2.mMaxResCnt AS MMaxResCnt, a2.mSuccess AS MSuccess").
		Joins("LEFT OUTER JOIN TblMaterialDrawIndex AS a2 ON a2.MDRD = a.MDRD").
		Joins("LEFT OUTER JOIN TblMaterialDrawMaterial AS a3 ON a3.MDID = a2.MDID").
		Joins("LEFT OUTER JOIN DT_Item AS b ON b.IID = a.IID").
		Joins("LEFT OUTER JOIN DT_Item AS b2 ON b2.IID = a3.IID").
		Limit(limit).
		Offset(offset).
		Scan(&result).Error; err != nil {
		return err
	}

	return c.JSON(result)
}

// Добавить Крафт
func PostAddCraft(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	refine := parm.Refine{
		RID:          1,
		RItemID0:     1,
		RItemID1:     1,
		RItemID2:     1,
		RItemID3:     1,
		RItemID4:     1,
		RItemID5:     1,
		RItemID6:     1,
		RItemID7:     1,
		RItemID8:     1,
		RItemID9:     1,
		RSuccess:     1,
		RIsCreateCnt: 1,
	}

	refineMaterial := parm.RefineMaterial{
		RID:      1,
		RItemID:  1,
		RNum:     1,
		ROrderNo: 1,
	}

	refineCreateInfo := parm.RefineCreateInfo{
		MIDX:      1,
		MRID:      1,
		MGroup1:   1,
		MGroup2:   1,
		MSort:     1,
		MItem0:    1,
		MItem1:    1,
		MItem2:    1,
		MItem3:    1,
		MCost:     1,
		MNationOp: 1,
	}

	fmt.Println(refine, refineMaterial, refineCreateInfo)

	return nil
}

/*
SET @ItemID = '123123' -- Номер создаваемого предмета
SET @Chance = '100' -- Шанс успешного создания
SET @Item1 = '1200' -- Первый предмет, участвующий в создании
SET @Item2 = '123122' -- Второй предмет, участвующий в создании
SET @Item3 = '1522' -- Третий предмет, участвующий в создании
SET @Item4 = '922' -- Четвертый предмет, участвующий в создании
SET @Cost = '0' -- Стоимость крафта в серебре
SET @Group1 = '3' -- Первая группа сортировки
SET @Group2 = '3' -- Вторая группа сортировки
SET @RID = (SELECT MAX(RID)+1 FROM [DT_Refine])
SET @NewIDX = (SELECT MAX(mIDX)+1 FROM [DT_RefineCreateInfo])
SET @Sort = (SELECT MAX(mSort)+1 FROM [DT_RefineCreateInfo] WHERE [mGroup1] = @Group1 AND [mGroup2] = @Group2)

INSERT INTO [dbo].[DT_Refine] ([RID], [RItemID0], [RItemID1], [RItemID2], [RItemID3], [RItemID4], [RItemID5], [RItemID6], [RItemID7], [RItemID8], [RItemID9], [RSuccess], [RIsCreateCnt])
VALUES
(@RID, @ItemID, 0, 0, 0, 0, 0, 0, 0, 0, 0, @Chance, 1);

INSERT INTO [dbo].[DT_RefineCreateInfo] ([mIDX], [mRID], [mGroup1], [mGroup2], [mSort], [mItem0], [mItem1], [mItem2], [mItem3], [mCost], [mNationOp])
VALUES
(@NewIDX, @RID, @Group1, @Group2, @Sort, 0, 0, 0, 0, @Cost, 1152921504606846975);

INSERT INTO [dbo].[DT_RefineMaterial] ([RID], [RItemID], [RNum], [ROrderNo]) VALUES (@RID, @Item1, 1, 1);
INSERT INTO [dbo].[DT_RefineMaterial] ([RID], [RItemID], [RNum], [ROrderNo]) VALUES (@RID, @Item2, 1, 2);
INSERT INTO [dbo].[DT_RefineMaterial] ([RID], [RItemID], [RNum], [ROrderNo]) VALUES (@RID, @Item3, 1, 3);
INSERT INTO [dbo].[DT_RefineMaterial] ([RID], [RItemID], [RNum], [ROrderNo]) VALUES (@RID, @Item4, 1, 4);
*/
