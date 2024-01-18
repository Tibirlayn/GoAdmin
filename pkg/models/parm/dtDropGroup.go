package parm

type DTdropGroup struct {
	DGroup	int `json:"MID" gorm:"column:MID"`
	DDrop	int `json:"DDrop" gorm:"column:DDrop"`
	DPercent float64 `json:"DPercent" gorm:"column:DPercent"`
}

func (DTdropGroup) TableName() string {
	return "DT_DropGroup"
}