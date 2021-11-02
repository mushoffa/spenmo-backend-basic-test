package model

import (
	"time"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type WalletDB struct {
	ID      string    `gorm:"id"`
	Created time.Time `gorm:"created"`
	Updated time.Time `gomr:"updated"`
	Name    string    `gorm:"name"`
	Balance float64   `gorm:"balance"`
	UserID  string    `gorm:"user_id"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (db *WalletDB) TableName() string {
	return "wallets"
}
