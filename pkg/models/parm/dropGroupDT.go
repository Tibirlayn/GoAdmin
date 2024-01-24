package parm

type DropGroupDT struct {
	DGroup   int         `json:"DGroup" gorm:"column:DGroup"`
	DDrop    int         `json:"DDrop" gorm:"column:DDrop"`
	DPercent float64     `json:"DPercent" gorm:"column:DPercent"`
}

func (DropGroupDT) TableName() string {
	return "DT_DropGroup"
}
