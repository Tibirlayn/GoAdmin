package billing

import "time"

type GoldItem struct {
	GoldItemID        int64     `json:"GoldItemID" gorm:"column:GoldItemID;not null"`
	IID               int       `json:"IID" gorm:"column:IID;not null"`
	ItemName          string    `json:"ItemName" gorm:"column:ItemName;not null;size:40"`
	ItemImage         string    `json:"ItemImage" gorm:"column:ItemImage;size:100"`
	ItemDesc          string    `json:"ItemDesc" gorm:"column:ItemDesc;size:500"`
	OriginalGoldPrice int       `json:"OriginalGoldPrice" gorm:"column:OriginalGoldPrice;not null"`
	GoldPrice         int       `json:"GoldPrice" gorm:"column:GoldPrice;not null"`
	ItemCategory      int16     `json:"ItemCategory" gorm:"column:ItemCategory;size:1;not null"`
	IsPackage         string    `json:"IsPackage" gorm:"column:IsPackage;size:1;not null"`
	Status            string    `json:"Status" gorm:"column:Status;size:1;not null"`
	AvailablePeriod   int       `json:"AvailablePeriod" gorm:"column:AvailablePeriod;not null"`
	Count             int       `json:"Count" gorm:"column:Count;not null"`
	PracticalPeriod   int       `json:"PracticalPeriod" gorm:"column:PracticalPeriod;not null"`
	RegistDate        time.Time `json:"RegistDate" gorm:"column:RegistDate;not null"`
	RegistAdmin       string    `json:"RegistAdmin" gorm:"column:RegistAdmin;size:20"`
	RegistIP          string    `json:"RegistIP" gorm:"column:RegistIP;size:19"`
	UpdateDate        string    `json:"UpdateDate" gorm:"column:UpdateDate"`
	UpdateAdmin       string    `json:"UpdateAdmin" gorm:"column:UpdateAdmin;size:20"`
	UpdateIP          string    `json:"UpdateIP" gorm:"column:UpdateIP;size:19"`
	ItemNameRUS       string    `json:"ItemNameRUS" gorm:"column:ItemNameRUS;size:255"`
	ItemDescRUS       string    `json:"ItemDescRUS" gorm:"column:ItemDescRUS;size:500"`
}

func (GoldItem) TableName() string {
	return "TBLGoldItem"
}
