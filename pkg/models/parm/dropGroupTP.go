package parm

type DropGroupTP struct {
	DGroup    int    `json:"DGroup" gorm:"column:DGroup;not null"`
	DName     string `json:"DName" gorm:"column:DName;size:40"`
	DDropType uint8  `json:"DDropType" gorm:"column:DDropType;not null"`
}

func (DropGroupTP) TableName() string {
	return "TP_DropGroup"
}
