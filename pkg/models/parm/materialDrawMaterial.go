package parm

type MaterialDrawMaterial struct {
	MSeq int   `json:"mSeq" gorm:"column:mSeq; not null"`
	MDID int64 `json:"MDID" gorm:"column:MDID; not null"`
	IID  int   `json:"IID" gorm:"column:IID; not null"`
	MCnt int   `json:"mCnt" gorm:"column:mCnt; not null"`
}

func (MaterialDrawMaterial) TableName() string {
	return "TblMaterialDrawMaterial"
}
