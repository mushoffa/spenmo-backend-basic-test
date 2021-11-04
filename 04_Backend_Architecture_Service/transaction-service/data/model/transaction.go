package model

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
type TransactionDB struct {
	ID              string  `gorm:"id"`
	Date            string  `gorm:"date"`
	Time            string  `gorm:"time"`
	AccountNumber   string  `gorm:"account_number"`
	AccountName     string  `gorm:"account_name"`
	PAN             string  `gorm:"pan"`
	Amount          float64 `gorm:"amount"`
	WalletID        string  `gorm:"wallet_id"`
	WalletName      string  `gorm:"wallet_name"`
	TransactionType string  `gorm:"transactionType"`
	ReferenceNumber string  `gorm:"reference_number"`
	Status          string  `gorm:"status"`
	Description     string  `gorm:"description"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (db *TransactionDB) TableName() string {
	return "transactions"
}
