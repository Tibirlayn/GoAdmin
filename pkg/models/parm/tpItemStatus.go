package parm

type TPitemStatus struct {
	MStatus	string `json:"mStatus" gorm:"column:mStatus;size:255"`
	MDesc	string `json:"mDesc" gorm:"column:mDesc;size:255"`
}

func (TPitemStatus) TableName() string {
	return "TP_ItemStatus"
}