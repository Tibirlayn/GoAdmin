package game

import (
	"time"
)

type Pc struct {
	MRegDate time.Time `json:"mRegDate" gorm:"column:mRegDate"`
	MOwner int `json:"mOwner" gorm:"column:mOwner"`
	MSlot int8 `json:"mSlot" gorm:"column:mSlot"`
	MNo	int `json:"mNo" gorm:"column:mNo;primaryKey"`
	MNm string `json:"mNm" gorm:"column:mNm;size:12"`
	MClass int8 `json:"mClass" gorm:"column:mClass"`
	MSex int8 `json:"mSex" gorm:"column:mSex"`
	MHead int8 `json:"mHead" gorm:"column:mHead"`
	MFace int8 `json:"mFace" gorm:"column:mFace"`
	MBody int8 `json:"mBody" gorm:"column:mBody"`
	MHomeMapNo int `json:"mHomeMapNo" gorm:"column:mHomeMapNo"`
	MHomePosX float64 `json:"mHomePosX" gorm:"column:mHomePosX"`
	MHomePosY float64 `json:"mHomePosY" gorm:"column:mHomePosY"`
	MHomePosZ float64 `json:"mHomePosZ" gorm:"column:mHomePosZ"`
	MDelDate time.Time `json:"mDelDate" gorm:"column:mDelDate"`
}

func (Pc) TableName() string {
	return "TblPc"
}