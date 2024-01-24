package parm

type MonsterDrop struct {
	MID      int         `json:"MID" gorm:"column:MID"`
	DGroup   int         `json:"DGroup" gorm:"column:DGroup"`
	DPercent int8        `json:"DPercent" gorm:"column:DPercent"`
}

func (MonsterDrop) TableName() string {
	return "DT_MonsterDrop"
}