package model

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
type LimitDB struct {
	ID           int64  `gorm:"id"`
	Rating       string `gorm:"rating"`
	LimitDaily   string `gorm:"limit_daily"`
	LimitMonthly string `gorm:"limit_monthly"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (db *LimitDB) TableName() string {
	return "limits"
}
