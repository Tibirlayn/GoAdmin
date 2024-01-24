package parm

type DropItem struct {
	DDrop	int `json:"DDrop" gorm:"column:DDrop"`
	DItem	int `json:"DItem" gorm:"column:DItem"`
	DNumber	int16 `json:"DNumber" gorm:"column:DNumber"`
	DStatus	uint8 `json:"DStatus" gorm:"column:DStatus;not null"`
	DIsEvent bool `json:"DIsEvent" gorm:"column:DIsEvent;not null"`
}

func (DropItem) TableName() string {
	return "DT_DropItem"
}