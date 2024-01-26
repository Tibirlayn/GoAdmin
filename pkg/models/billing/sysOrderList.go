package billing

import "time"

type SysOrderList struct {
	MRegDate         time.Time `json:"mRegDate" gorm:"mRegDate;not null"`
	MSysOrderID      int64     `json:"mSysOrderID" gorm:"mSysOrderID;not null"`
	MSysID           int64     `json:"mSysID" gorm:"mSysID;not null"`
	MUserNo          int       `json:"mUserNo" gorm:"mUserNo;not null"`
	MSvrNo           int16     `json:"mSvrNo" gorm:"mSvrNo;not null"`
	MItemID          int       `json:"mItemID" gorm:"mItemID;not null"`
	MCnt             int       `json:"mCnt" gorm:"mCnt;not null"`
	MAvailablePeriod int       `json:"mAvailablePeriod" gorm:"mAvailablePeriod"`
	MPracticalPeriod int       `json:"mPracticalPeriod" gorm:"mPracticalPeriod"`
	MStatus          uint8     `json:"mStatus" gorm:"mStatus;not null"`
	MReceiptDate     time.Time `json:"mReceiptDate" gorm:"mReceiptDate"`
	MReceiptPcNo     int       `json:"mReceiptPcNo" gorm:"mReceiptPcNo"`
	MRecepitPcNm     string    `json:"mRecepitPcNm" gorm:"mRecepitPcNm"`
	MBindingType     uint8     `json:"mBindingType" gorm:"mBindingType"`
	MLimitedDate     time.Time `json:"mLimitedDate" gorm:"mLimitedDate"`
	MItemStatus      uint8     `json:"mItemStatus" gorm:"mItemStatus"`
}
