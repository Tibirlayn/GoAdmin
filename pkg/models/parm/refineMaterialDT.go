package parm

type RefineMaterial struct {
	RID      int  `json:"RID" gorm:"column:RID"`
	RItemID  int  `json:"RItemID" gorm:"column:RItemID"`
	RNum     int  `json:"RNum" gorm:"column:RNum"`
	ROrderNo int8 `json:"ROrderNo" gorm:"column:ROrderNo;not null"`
}

func (RefineMaterial) TableName() string {
	return "DT_RefineMaterial"
}
