package controllers

import (
	"github.com/Tibirlayn/GoAdmin/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func GetInfoBossDrop(c *fiber.Ctx) error {

	if DB, err := config.ParmConfiguration(); err != nil {
		return err
	} else {
		result := DB.Table("DT_Monster").
        Select("mo.MName AS Boss, tdg.DName AS 'Name group', it.IName AS 'Name item', tpis.mDesc AS 'Item', CASE WHEN evo.mObjID IS NULL THEN 'Not event item' ELSE 'Event item' END AS 'Event'").
        Joins("INNER JOIN DT_MonsterDrop AS md ON mo.MID = md.MID").
        Joins("INNER JOIN DT_DropGroup AS dg ON md.DGroup = dg.DGroup").
        Joins("INNER JOIN TP_DropGroup AS tdg ON md.DGroup = tdg.DGroup").
        Joins("INNER JOIN DT_DropItem AS di ON dg.DDrop = di.DDrop").
        Joins("INNER JOIN TP_ItemStatus AS tpis ON di.DStatus = tpis.mStatus").
        Joins("INNER JOIN DT_Item AS it ON di.DItem = it.IID").
        Joins("LEFT JOIN TblEventObj AS evo ON it.IID = evo.mObjID").
        Where("mo.MID = ? and di.DIsEvent = 0", 84).
        Find(&bossDrops).Error
	}

	




    return c.JSON()
}

