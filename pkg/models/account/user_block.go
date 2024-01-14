package account

import "time"

type UserBlock struct{
	MRegDate time.Time `json:"mRegDate" gorm:"column:mRegDate;not null"`
	MUserNo int `json:"mUserNo" gorm:"column:mUserNo;not null;primaryKey"`
	MCertify time.Time `json:"mCertify" gorm:"column:mCertify"`
	MBoard time.Time `json:"mBoard" gorm:"column:mBoard"`
	MChat int `json:"mChat" gorm:"column:mChat"`
	MCertifyReason string `json:"mCertifyReason" gorm:"column:mCertifyReason;size:200"`
}

func (UserBlock) TableName() string {
	return "TblUserBlock"
}