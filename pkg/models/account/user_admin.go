package account

type UserAdmin struct {
	MUserNo int    `json:"mUserNo" gorm:"primaryKey;column:mUserNo"`
	MUserId string `json:"mUserId" gorm:"size:20;column:mUserId"`
}

func (UserAdmin) TableName() string {
	return "TblUserAdmin"
}

