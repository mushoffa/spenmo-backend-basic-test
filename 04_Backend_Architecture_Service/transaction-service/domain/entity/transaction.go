package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated 04/11/2021
type Transaction struct {
	ID              string  `json:"id"`
	Date            string  `json:"date"`
	Time            string  `json:"time"`
	AccountNumber   string  `json:"account_number"`
	AccountName     string  `json:"account_name"`
	PAN             string  `json:"pan"`
	Amount          float64 `json:"amount"`
	WalletID        string  `json:"wallet_id"`
	WalletName      string  `json:"wallet_name"`
	TransactionType string  `json:"transactionType"`
	ReferenceNumber string  `json:"reference_number"`
	Status          string  `json:"status"`
	Description     string  `json:"description"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 04/11/2021
// @Updated
func (e *Transaction) GenerateID() {
	e.ID = uuid.New().String()
}

// @Author Ahmad Ridwan Mushoffa
// @Created 04/11/2021
// @Updated
func (e *Transaction) GenerateDateTime() {
	today := time.Now()

	dateLayout := "02-01-2006"
	timeLayout := "15:04:05"

	e.Date = today.Format(dateLayout)
	e.Time = today.Format(timeLayout)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 04/11/2021
// @Updated
func (e *Transaction) GenerateReferenceNumber(transactionReference string) {
	timestamp := time.Now().Unix()

	e.ReferenceNumber = fmt.Sprintf("%s-%s-%d", e.TransactionType, transactionReference, timestamp)
}