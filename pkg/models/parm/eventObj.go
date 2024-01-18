package parm

type EventObj struct {
	MID			int `json:"mID" gorm:"column:mID;not null"`
	MEventID	int `json:"mEventID" gorm:"column:mEventID;not null"`
	MObjType	uint8 `json:"mObjType" gorm:"column:mObjType;not null"`
	MObjID		int `json:"mObjID" gorm:"column:mObjID;not null"`
}

func (EventObj) TableName() string {
	return "TblEventObj"
}