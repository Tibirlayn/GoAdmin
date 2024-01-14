package account

type Member struct {
    MUserId      string  `json:"mUserId" gorm:"column:mUserId;size:50;not null;primaryKey"`
    MUserPswd    string  `json:"mUserPswd" gorm:"column:mUserPswd;size:50"`
    Superpwd     string  `json:"superpwd" gorm:"column:superpwd;size:50"`
    Cash         float64 `json:"cash" gorm:"column:cash"`
    Email        string  `json:"email" gorm:"column:email;size:255"`
    Tgzh         string  `json:"tgzh" gorm:"column:tgzh;size:255"`
    Uid          int     `json:"uid" gorm:"column:uid;not null"`
    Klq          int     `json:"klq" gorm:"column:klq"`
    Ylq          int     `json:"ylq" gorm:"column:ylq"`
    Auth         int     `json:"auth" gorm:"column:auth"`
    MSum         string  `json:"mSum" gorm:"column:mSum;size:255"`
    IsAdmin      int     `json:"isAdmin" gorm:"column:isAdmin"`
    Isdl         int     `json:"isdl" gorm:"column:isdl"`
    Dlmoney      int     `json:"dlmoney" gorm:"column:dlmoney"`
    RegisterIp   string  `json:"registerIp" gorm:"column:registerIp;size:255"`
    Country      string  `json:"country" gorm:"column:country;size:20"`
    CashBack     int     `json:"cashBack" gorm:"column:cashBack"`
}

func (Member) TableName() string {
    return "Member"
}