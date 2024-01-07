package account

import "time"

type TblUser struct {
	MRegDate           time.Time `json:"mRegDate"`
	MUserAuth          int8      `json:"mUserAuth"`
	MUserNo            int       `json:"mUserNo"`
	MUserId            string    `json:"mUserId"`
	MUserPswd          string    `json:"mUserPswd"`
	MCertifiedKey      int       `json:"mCertifiedKey,omitempty"`
	MIp                string    `json:"mIp"`
	MLoginTm           time.Time `json:"mLoginTm,omitempty"`
	MLogoutTm          time.Time `json:"mLogoutTm,omitempty"`
	MTotUseTm          int       `json:"mTotUseTm"`
	MWorldNo           int16     `json:"mWorldNo"`
	MDelDate           time.Time `json:"mDelDate"`
	MPcBangLv          int       `json:"mPcBangLv"`
	MSecKeyTableUse    int8      `json:"mSecKeyTableUse"`
	MUseMacro          int16     `json:"mUseMacro"`
	MIpEX              int64     `json:"mIpEX"`
	MJoinCode          string    `json:"mJoinCode"`
	MLoginChannelID    string    `json:"mLoginChannelID"`
	MTired             string    `json:"mTired"`
	MChnSID            string    `json:"mChnSID"`
	MNewId             bool      `json:"mNewId"`
	MLoginSvrType      int8      `json:"mLoginSvrType"`
	MAccountGuid       int       `json:"mAccountGuid"`
	MNormalLimitTime   int       `json:"mNormalLimitTime"`
	MPcBangLimitTime   int       `json:"mPcBangLimitTime"`
	MRegIp             string    `json:"mRegIp,omitempty"`
	MIsMovingToBattleSvr bool    `json:"mIsMovingToBattleSvr,omitempty"`
	
}
