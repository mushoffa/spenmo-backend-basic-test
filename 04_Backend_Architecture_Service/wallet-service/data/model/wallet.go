package model

import (
	"time"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 03/11/2021
type WalletDB struct {
	ID            string    `gorm:"id"`
	Created       time.Time `gorm:"created"`
	Updated       time.Time `gorm:"updated"`
	Name          string    `gorm:"name"`
	Balance       float64   `gorm:"balance"`
	MaxLimit      float64   `gorm:"max_limit"`
	AccountNumber string    `gorm:"account_number"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (db *WalletDB) TableName() string {
	return "wallets"
}
