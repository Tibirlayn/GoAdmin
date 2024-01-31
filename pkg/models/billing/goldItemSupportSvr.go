package billing

type GoldItemSupportSvr struct {
	GoldItemID int64 `json:"GoldItemID" gorm:"colomn:GoldItemID;not null"`
	MSvrNo     int16 `json:"mSvrNo" gorm:"colomn:mSvrNo;not null"`
}

func (GoldItemSupportSvr) TableName() string {
	return "TBLGoldItemSupportSvr"
}
