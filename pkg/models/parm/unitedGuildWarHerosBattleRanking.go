package parm

import "time"

type UnitedGuildWarHerosBattleRanking struct {
	MRegDate    time.Time `json:"mRegDate" gorm:"column:mRegDate; not null"`
	MSvrNo      int16     `json:"mSvrNo" gorm:"column:mSvrNo; not null"`
	MRanking    int16     `json:"mRanking" gorm:"column:mRanking; not null"`
	MGuildNo    int       `json:"mGuildNo" gorm:"column:mGuildNo; not null"`
	MGuildNm    string    `json:"mGuildNm" gorm:"column:mGuildNm; not null; size:12"`
	MVictoryCnt int       `json:"mVictoryCnt" gorm:"column:mVictoryCnt; not null"`
	MPcNo       int       `json:"mPcNo" gorm:"column:mPcNo; not null"`
	MPcNm       string    `json:"mPcNm" gorm:"column:mPcNm; not null; size:12"`
}

func (UnitedGuildWarHerosBattleRanking) TableName() string {
	return "TblUnitedGuildWarHerosBattleRanking"
}
