package parm

type RefineCreateInfo struct {
	MIDX	int `json:"mIDX" gorm:"column:mIDX; not null"`
	MRID	int `json:"mRID" gorm:"column:mRID; not null"`
	MGroup1	int8 `json:"mGroup1" gorm:"column:mGroup1; not null"`
	MGroup2	int8 `json:"mGroup2" gorm:"column:mGroup2; not null"`
	MSort	int `json:"mSort" gorm:"column:mSort; not null"`
	MItem0	int `json:"mItem0" gorm:"column:mItem0; not null"`
	MItem1	int `json:"mItem1" gorm:"column:mItem1; not null"`
	MItem2	int `json:"mItem2" gorm:"column:mItem2; not null"`
	MItem3	int `json:"mItem3" gorm:"column:mItem3; not null"`
	MCost	int `json:"mCost" gorm:"column:mCost; not null"`
	MNationOp	int64 `json:"mNationOp" gorm:"column:mNationOp; not null"`
}

func (RefineCreateInfo) TableName() string {
	return "DT_RefineCreateInfo"
}
