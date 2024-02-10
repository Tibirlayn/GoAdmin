package parm

type MaterialDrawResult struct {
	MSeq        int     `json:"mSeq" gorm:"column:mSeq; not null"`
	MDRD        int64   `json:"MDRD" gorm:"column:MDRD; not null"`
	IID         int     `json:"IID" gorm:"column:IID; not null"`
	MPerOrRate  float64 `json:"mPerOrRate" gorm:"column:mPerOrRate"`
	MItemStatus int8    `json:"mItemStatus" gorm:"column:mItemStatus; not null"`
	MCnt        int     `json:"mCnt" gorm:"column:mCnt; not null"`
	MBinding    int     `json:"mBinding" gorm:"column:mBinding; not null"`
	MEffTime    int     `json:"mEffTime" gorm:"column:mEffTime; not null"`
	MValTime    int16   `json:"mValTime" gorm:"column:mValTime; not null"`
	MResource   int     `json:"mResource" gorm:"column:mResource; not null"`
	MAddGroup   int8    `json:"mAddGroup" gorm:"column:mAddGroup; not null"`
}

func (MaterialDrawResult) TableName() string {
	return "TblMaterialDrawResult"
}
