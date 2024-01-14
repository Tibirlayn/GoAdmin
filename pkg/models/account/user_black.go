package account

import "time"

type UserBlack struct {
	MRegDate time.Time `json:"mRegDate" gorm:"column:mRegDate;not null"`
	MUserId string `json:"mUserId" gorm:"column:mUserId;not null;size:20;primaryKey"`
	MDesc string `json:"mDesc" gorm:"column:mDesc;size:200"`
}

func (UserBlack) TableName() string {
	return "TblUserBlack"
}
