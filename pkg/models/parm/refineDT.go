package parm

type Refine struct {
	RID          int     `json:"RID" gorm:"column:RID;not null;primaryKey"`
	RItemID0     int     `json:"RItemID0" gorm:"column:RItemID0;not null"`
	RItemID1     int     `json:"RItemID1" gorm:"column:RItemID1;not null"`
	RItemID2     int     `json:"RItemID2" gorm:"column:RItemID2;not null"`
	RItemID3     int     `json:"RItemID3" gorm:"column:RItemID3;not null"`
	RItemID4     int     `json:"RItemID4" gorm:"column:RItemID4;not null"`
	RItemID5     int     `json:"RItemID5" gorm:"column:RItemID5;not null"`
	RItemID6     int     `json:"RItemID6" gorm:"column:RItemID6;not null"`
	RItemID7     int     `json:"RItemID7" gorm:"column:RItemID7;not null"`
	RItemID8     int     `json:"RItemID8" gorm:"column:RItemID8;not null"`
	RItemID9     int     `json:"RItemID9" gorm:"column:RItemID9;not null"`
	RSuccess     float64 `json:"RSuccess" gorm:"column:RSuccess"`
	RIsCreateCnt int16   `json:"RIsCreateCnt" gorm:"column:RIsCreateCnt;not null"`
}

func (Refine) TableName() string {
	return "DT_Refine"
}