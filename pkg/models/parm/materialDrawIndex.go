package parm

type MaterialDrawIndex struct {
	MDID             int64   `json:"MDID" gorm:"column:MDID; not null"`
	MDRD             int64   `json:"MDRD" gorm:"column:MDRD; not null"`
	MResType         int     `json:"mResType" gorm:"column:mResType; not null"`
	MMaxResCnt       int     `json:"mMaxResCnt" gorm:"column:mMaxResCnt; not null"`
	MSuccess         float64 `json:"mSuccess" gorm:"column:mSuccess"`
	MDesc            string  `json:"mDesc" gorm:"column:mDesc; not null; size:500"`
	MAddQuestionMark int16   `json:"mAddQuestionMark" gorm:"column:mAddQuestionMark; not null"`
	MDescRus         string  `json:"mDescRus" gorm:"column:mDescRus; size:2000"`
}

func (MaterialDrawIndex) TableName() string {
	return "TblMaterialDrawIndex"
}
