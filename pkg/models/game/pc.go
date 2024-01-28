package game

import (
	"time"
)

type Pc struct {
	MRegDate time.Time `json:"mRegDate" gorm:"column:mRegDate"` 
	MOwner int `json:"mOwner" gorm:"column:mOwner"` // аккаунт персонажа
	MSlot int8 `json:"mSlot" gorm:"column:mSlot"` // слот персонажа в меню выбара [0, 1, 2] 
	MNo	int `json:"mNo" gorm:"column:mNo;primaryKey"` // id персонажа
	MNm string `json:"mNm" gorm:"column:mNm;size:12"` // имя персонажа
	MClass int8 `json:"mClass" gorm:"column:mClass"` // класс персонажа 
	MSex int8 `json:"mSex" gorm:"column:mSex"` // пол персонажа
	MHead int8 `json:"mHead" gorm:"column:mHead"` // волосы персонажа 
	MFace int8 `json:"mFace" gorm:"column:mFace"` // лицо персонажа 
	MBody int8 `json:"mBody" gorm:"column:mBody"` // тело 
	MHomeMapNo int `json:"mHomeMapNo" gorm:"column:mHomeMapNo"` 
	MHomePosX float64 `json:"mHomePosX" gorm:"column:mHomePosX"` // координаты расположения персонажа
	MHomePosY float64 `json:"mHomePosY" gorm:"column:mHomePosY"` // координаты расположения персонажа
	MHomePosZ float64 `json:"mHomePosZ" gorm:"column:mHomePosZ"` // координаты расположения персонажа
	MDelDate time.Time `json:"mDelDate" gorm:"column:mDelDate"` // если стоит дата персонаж удален
}

func (Pc) TableName() string {
	return "TblPc"
}