package parm

type Moster struct {
	MID	int `json:"MID" gorm:"column:MID;primaryKey;not null"`
	MName string `json:"MName" gorm:"column:MName;size:40"`
	MLevel	int `json:"mLevel" gorm:"column:mLevel"`
	MClass	int `json:"MClass" gorm:"column:MClass"`
	MExp	int `json:"MExp" gorm:"column:MExp"`
	MHIT	int16  `json:"MHIT" gorm:"column:MHIT"`
	MMinD	int16 `json:"MMinD" gorm:"column:MMinD"`
	MMaxD	int16 `json:"MMaxD" gorm:"column:MMaxD"`
	MAttackRateOrg	int16 `json:"MAttackRateOrg" gorm:"column:MAttackRateOrg"`
	MMoveRateOrg	int16 `json:"MMoveRateOrg" gorm:"column:MMoveRateOrg"`
	MAttackRateNew	int16 `json:"MAttackRateNew" gorm:"column:MAttackRateNew"`
	MMoveRateNew	int16 `json:"MMoveRateNew" gorm:"column:MMoveRateNew"`
	MHP	int16 `json:"MHP" gorm:"column:MHP"`
	MMP	int16 `json:"MMP" gorm:"column:MMP"`
	MMoveRange	int16 `json:"MMoveRange" gorm:"column:MMoveRange;not null"`
	MGbjType	int16 `json:"MGbjType" gorm:"column:MGbjType;not null"`
	MRaceType	int16 `json:"MRaceType" gorm:"column:MRaceType"`
	MAiType	int16 `json:"MAiType" gorm:"column:MAiType;not null"`
	MCastingDelay	int16 `json:"MCastingDelay" gorm:"column:MCastingDelay;not null"`
	MChaotic	int16 `json:"MChaotic" gorm:"column:MChaotic;not null"`
	MSameRace1	int `json:"MSameRace1" gorm:"column:MSameRace1;not null"`
	MSameRace2	int `json:"MSameRace2" gorm:"column:MSameRace2;not null"`
	MSameRace3	int `json:"MSameRace3" gorm:"column:MSameRace3;not null"`
	MSameRace4	int `json:"MSameRace4" gorm:"column:MSameRace4;not null"`
	MSightRange	int `json:"mSightRange" gorm:"column:mSightRange;not null"`
	MAttackRange	int `json:"mAttackRange" gorm:"column:mAttackRange;not null"`
	MSkillRange	int `json:"mSkillRange" gorm:"column:mSkillRange;not null"`
	MBodySize	int `json:"mBodySize" gorm:"column:mBodySize;not null"`
	MDetectTransF	int16 `json:"mDetectTransF" gorm:"column:mDetectTransF;not null"`
	MDetectTransP	int16 `json:"mDetectTransP" gorm:"column:mDetectTransP;not null"`
	MDetectChao	int16 `json:"mDetectChao" gorm:"column:mDetectChao;not null"`
	MAiEx	int `json:"mAiEx" gorm:"column:mAiEx"`
	MScale	float64 `json:"mScale" gorm:"column:mScale;not null;size:53"`
	MIsResistTransF	bool `json:"mIsResistTransF" gorm:"column:mIsResistTransF;not null"`
	MIsEvent	bool `json:"mIsEvent" gorm:"column:mIsEvent;not null"`
	MIsTest	bool `json:"mIsTest" gorm:"column:mIsTest;not null"`
	MHPNew	int16 `json:"mHPNew" gorm:"column:mHPNew;not null"`
	MMPNew	int16 `json:"mMPNew" gorm:"column:mMPNew;not null"`
	MBuyMerchanID	int `json:"mBuyMerchanID" gorm:"column:mBuyMerchanID;not null"`
	MSellMerchanID	int `json:"mSellMerchanID" gorm:"column:mSellMerchanID;not null"`
	MChargeMerchanID	int `json:"mChargeMerchanID" gorm:"column:mChargeMerchanID;not null"`
	MTransformWeight	int16 `json:"mTransformWeight" gorm:"column:mTransformWeight;not null"`
	MNationOp	int64 `json:"mNationOp" gorm:"column:mNationOp;not null"`
	MHPRegen	int16 `json:"mHPRegen" gorm:"column:mHPRegen;not null"`
	MMPRegen	int16 `json:"mMPRegen" gorm:"column:mMPRegen;not null"`
	IContentsLv	uint8 `json:"IContentsLv" gorm:"column:IContentsLv;not null"`
	MIsEventTest	bool `json:"mIsEventTest" gorm:"column:mIsEventTest;not null"`
	MIsShowHp	bool `json:"mIsShowHp" gorm:"column:mIsShowHp;not null"`
	MSupportType	uint8 `json:"mSupportType" gorm:"column:mSupportType;not null"`
	MVolitionOfHonor	int16 `json:"mVolitionOfHonor" gorm:"column:mVolitionOfHonor;not null"`
	MWMapIconType	int16 `json:"mWMapIconType" gorm:"column:mWMapIconType;not null"`
	MIsAmpliableTermOfValidity	bool `json:"mIsAmpliableTermOfValidity" gorm:"column:mIsAmpliableTermOfValidity;not null"`
	MAttackType	uint8 `json:"mAttackType" gorm:"column:mAttackType;not null"`
	MTransType	uint8 `json:"mTransType" gorm:"column:mTransType;not null"`
	MDPV	int16 `json:"mDPV" gorm:"column:mDPV;not null"`
	MMPV	int16 `json:"mMPV" gorm:"column:mMPV;not null"`
	MRPV	int16 `json:"mRPV" gorm:"column:mRPV;not null"`
	MDDV	int16 `json:"mDDV" gorm:"column:mDDV;not null"`
	MMDV	int16 `json:"mMDV" gorm:"column:mMDV;not null"`
	MRDV	int16 `json:"mRDV" gorm:"column:mRDV;not null"`
	MSubDDWhenCritical	int16 `json:"mSubDDWhenCritical" gorm:"column:mSubDDWhenCritical;not null"`
	MEnemySubCriticalHit	int16 `json:"mEnemySubCriticalHit" gorm:"column:mEnemySubCriticalHit;not null"`
	MEventQuest	uint8 `json:"mEventQuest" gorm:"column:mEventQuest;not null"`
	MEScale	float64 `json:"mEScale" gorm:"column:mEScale;not null;size:53"`
}

func (Moster) TableName() string {
	return "DT_Monster"
}