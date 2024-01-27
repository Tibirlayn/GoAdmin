package billing

import (
	"time"
)

type SysOrderList struct {
	MRegDate         time.Time `json:"mRegDate" gorm:"column:mRegDate;not null"`
	MSysOrderID      int64     `json:"mSysOrderID" gorm:"column:mSysOrderID;not null;autoIncrement"`
	MSysID           int64     `json:"mSysID" gorm:"column:mSysID;not null"` // BIGINT = 196491, /* — ID сообщения от администратора */
	MUserNo          int       `json:"mUserNo" gorm:"column:mUserNo;not null"` // id аккаунт
	MSvrNo           int16     `json:"mSvrNo" gorm:"column:mSvrNo;not null"` // SMALLINT = 9991, /* — Номер вашего сервера */
	MItemID          int       `json:"mItemID" gorm:"column:mItemID;not null"` // INT = 409, /* — Номер предмета (подарка) */
	MCnt             int       `json:"mCnt" gorm:"column:mCnt;not null"` // INT = 1000, /* — Количество */
	MAvailablePeriod int       `json:"mAvailablePeriod" gorm:"column:mAvailablePeriod"` // INT = 0, /* — Доступный период (сколько будет лежать в подароках) */
	MPracticalPeriod int       `json:"mPracticalPeriod" gorm:"column:mPracticalPeriod"` // INT = 0, /* — Практический период (количество времени которое будет у предмета после получения)*/
	MStatus          uint8     `json:"mStatus" gorm:"column:mStatus;not null"`
	MReceiptDate     time.Time `json:"mReceiptDate" gorm:"column:mReceiptDate"`
	MReceiptPcNo     int       `json:"mReceiptPcNo" gorm:"column:mReceiptPcNo"` // персонаж
	MRecepitPcNm     string    `json:"mRecepitPcNm" gorm:"column:mRecepitPcNm"`
	MBindingType     uint8     `json:"mBindingType" gorm:"column:mBindingType"` // TINYINT = 0, /* — Под замком предмет или нет (Нет = 0 | Да = 1) */
	MLimitedDate     time.Time `json:"mLimitedDate" gorm:"column:mLimitedDate"` // SMALLDATETIME = '2079-06-06', /* — Ограниченная дата */
	MItemStatus      uint8     `json:"mItemStatus" gorm:"column:mItemStatus"` // TINYINT = 1; /* — Статус предмета */
}

func (SysOrderList) TableName() string {
	return "TBLSysOrderList"
}