package model

import "time"

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type CardDB struct {
	ID           string    `gorm:"id"`
	Created      time.Time `gorm:"created"`
	Updated      time.Time `gorm:"updated"`
	PAN          string    `gorm:"pan"`
	Name         string    `gorm:"name"`
	ExpiryDate   string    `gorm:"expiry_date"`
	LimitDaily   float64   `gorm:"limit_daily"`
	LimitMonthly float64   `gorm:"limit_monthly"`
	CardType     string    `gorm:"card_type"`
	Status       string    `gorm:"status"`
	PIN          string    `gorm:"pin"`
	WalletID     string    `gorm:"wallet_id"`
	UserID       string    `gorm:"user_id"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (db *CardDB) TableName() string {
	return "cards"
}
