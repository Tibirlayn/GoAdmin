package billing

import "time"

type CategoryAssign struct {
	CategoryID  int16     `json:"CategoryID" gorm:"colomn:CategoryID;not null"`            // smallint
	GoldItemID  int64     `json:"GoldItemID" gorm:"colomn:GoldItemID;not null;primaryKey"` // bigint
	Status      string    `json:"Status" gorm:"colomn:Status;not null"`                    // nchar	1
	OrderNO     int16     `json:"OrderNO" gorm:"colomn:OrderNO;not null"`                  // smallint
	RegistDate  time.Time `json:"RegistDate" gorm:"colomn:RegistDate;not null"`            // datetime
	RegistAdmin string    `json:"RegistAdmin" gorm:"colomn:RegistAdmin"`                   // nvarchar	20
	RegistIP    string    `json:"RegistIP" gorm:"colomn:RegistIP"`                         // nvarchar	19
	UpdateDate  time.Time `json:"UpdateDate" gorm:"colomn:UpdateDate"`                     // datetime
	UpdateAdmin string    `json:"UpdateAdmin" gorm:"colomn:UpdateAdmin"`                   // nvarchar	20
	UpdateIP    string    `json:"UpdateIP" gorm:"colomn:UpdateIP"`                         // nvarchar	19

}

func (CategoryAssign) TableName() string {
	return "TBLCategoryAssign"
}
