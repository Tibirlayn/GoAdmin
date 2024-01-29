package parm

type ItemResource struct {
	RID       int    `json:"RID" gorm:"column:RID;not null;primaryKey"`
	ROwnerID  int    `json:"ROwnerID" gorm:"column:ROwnerID"`
	RType     int    `json:"RType" gorm:"column:RType"`
	RFileName string `json:"RFileName" gorm:"column:RFileName"`
	RPosX     int    `json:"RPosX" gorm:"column:RPosX"`
	RPosY     int    `json:"RPosY" gorm:"column:RPosY"`
}

func (ItemResource) TableName() string {
	return "DT_ItemResource"
}