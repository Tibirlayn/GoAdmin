package controllers

import (
	"errors"
	"fmt"
	"strconv"

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

	var id int
	success, _ := strconv.ParseFloat(data["Success"], 64)
	ItemID, _ := strconv.Atoi(data["ItemID"])

	ParmDB, err := config.ParmConfiguration()
	if err != nil {
		return err
	}

	if err := ParmDB.Table("DT_Refine").Select("MAX(RID) + 1").Scan(&id).Error; err != nil {
		return err
	}

	refine := parm.Refine{
		RID:          id,
		RItemID0:     ItemID, // Номер создаваемого предмета
		RItemID1:     0,
		RItemID2:     0,
		RItemID3:     0,
		RItemID4:     0,
		RItemID5:     0,
		RItemID6:     0,
		RItemID7:     0,
		RItemID8:     0,
		RItemID9:     0,
		RSuccess:     success, // Шанс успешного создания
		RIsCreateCnt: 1,
	}

	tx := ParmDB.Begin()
	if err := ParmDB.Create(&refine).Error; err != nil {
		tx.Rollback() // Откатить транзакцию при возникновении ошибки
		return err
	}

	var idx, sort int
	group1, _ := strconv.ParseInt(data["Group1"], 10, 8)
	group2, _ := strconv.ParseInt(data["Group2"], 10, 8)
	cost, _ := strconv.Atoi(data["Cost"])
	if err := ParmDB.Table("DT_RefineCreateInfo").
		Select("MAX(mIDX) + 1").
		Scan(&idx).Error; err != nil {
		return err
	}

	if err := ParmDB.Table("DT_RefineCreateInfo").
		Select("MAX(mSort) + 1").
		Where("mGroup1 = ? AND mGroup2 = ?").
		Scan(&sort).Error; err != nil {
		return err
	}

	refineCreateInfo := parm.RefineCreateInfo{
		MIDX:      idx,
		MRID:      id,
		MGroup1:   int8(group1),
		MGroup2:   int8(group2),
		MSort:     sort,
		MItem0:    0,
		MItem1:    0,
		MItem2:    0,
		MItem3:    0,
		MCost:     cost,
		MNationOp: 1152921504606846975,
	}

	if err := ParmDB.Create(&refineCreateInfo).Error; err != nil {
		tx.Rollback()
		return err
	}

	var items []int
	item1, _ := strconv.Atoi(data["Item1"])
	item2, _ := strconv.Atoi(data["Item2"])
	item3, _ := strconv.Atoi(data["Item3"])
	item4, _ := strconv.Atoi(data["Item4"])

	items = append(items, item1, item2, item3, item4)

	for key, item := range items {
		RefineMaterial := parm.RefineMaterial{
			RID:      id,
			RItemID:  item,
			RNum:     1,
			ROrderNo: int8(key) + 1,
		}

		if err := ParmDB.Create(RefineMaterial).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

// // Посмотреть дроп из золотого/изумрудного сундука + шансы
func GetGoldChest(c *fiber.Ctx) error {
	ParmDB, err := config.ParmConfiguration()
	if err != nil {
		return err
	}

	var goldenChest = 929
	var emeraldChest = 2578
	var resultGolden string
	var resultEmerald string

	if err := ParmDB.Table("TblDialogScript").Select("mScriptText").Where("mMId = ?", goldenChest).Scan(&resultGolden).Error; err != nil {
		return err
	}

	if err := ParmDB.Table("TblDialogScript").Select("mScriptText").Where("mMId = ?", emeraldChest).Scan(&resultEmerald).Error; err != nil {
		return err
	}

	return nil
}
