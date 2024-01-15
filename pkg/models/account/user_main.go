package account

import "time"

type TblUser struct {
	MRegDate           time.Time `json:"mRegDate" gorm:"column:mRegDate"`
	MUserAuth          uint8     `json:"mUserAuth" gorm:"column:mUserAuth"`
	MUserNo            int       `json:"mUserNo" gorm:"column:mUserNo;primaryKey"`
	MUserId            string    `json:"mUserId" gorm:"column:mUserId"`
	MUserPswd          string    `json:"mUserPswd" gorm:"column:mUserPswd"`
	MCertifiedKey      int       `json:"mCertifiedKey" gorm:"column:mCertifiedKey"`
	MIp                string    `json:"mIp" gorm:"column:mIp"`
	MLoginTm           time.Time `json:"mLoginTm" gorm:"not null;column:mLoginTm"`
	MLogoutTm          time.Time `json:"mLogoutTm" gorm:"not null;column:mLogoutTm"`
	MTotUseTm          int       `json:"mTotUseTm" gorm:"column:mTotUseTm"`
	MWorldNo           int16     `json:"mWorldNo" gorm:"column:mWorldNo"`
	MDelDate           time.Time `json:"mDelDate" gorm:"column:mDelDate"`
	MPcBangLv          int       `json:"mPcBangLv" gorm:"column:mPcBangLv"`
	MSecKeyTableUse    uint8     `json:"mSecKeyTableUse" gorm:"column:mSecKeyTableUse"`
	MUseMacro          int16     `json:"mUseMacro" gorm:"column:mUseMacro"`
	MIpEX              int64     `json:"mIpEx" gorm:"column:mIpEx"`
	MJoinCode          string    `json:"mJoinCode" gorm:"column:mJoinCode"`
	MLoginChannelID    string    `json:"mLoginChannelID" gorm:"column:mLoginChannelID"`
	MTired             string    `json:"mTired" gorm:"column:mTired"`
	MChnSID            string    `json:"mChnSID" gorm:"column:mChnSID"`
	MNewId             bool      `json:"mNewId" gorm:"column:mNewId"`
	MLoginSvrType      uint8     `json:"mLoginSvrType" gorm:"column:mLoginSvrType"`
	MAccountGuid       int       `json:"mAccountGuid" gorm:"column:mAccountGuid"`
	MNormalLimitTime   int       `json:"mNormalLimitTime" gorm:"column:mNormalLimitTime"`
	MPcBangLimitTime   int       `json:"mPcBangLimitTime" gorm:"column:mPcBangLimitTime"`
	MRegIp             string    `json:"mRegIp" gorm:"not null;column:mRegIp"`
	MIsMovingToBattleSvr bool    `json:"mIsMovingToBattleSvr" gorm:"not null;column:mIsMovingToBattleSvr"`
}


func (TblUser) TableName() string {
	return "TblUser"
}
