package parm

type MonsterResource struct {
	RID       int    `json:"RID" gorm:"column:RID;not null"`
	ROwnerID  int    `json:"ROwnerID" gorm:"column:ROwnerID"`
	RType     int    `json:"RType" gorm:"column:RType"`
	RFileName string `json:"RFileName" gorm:"column:RFileName;size:50"`
	RPosX     int    `json:"RPosX" gorm:"column:RPosX"`
	RPosY     int    `json:"RPosY" gorm:"column:RPosY"`
}

func (MonsterResource) TableName() string {
	return "DT_MonsterResource"
}
