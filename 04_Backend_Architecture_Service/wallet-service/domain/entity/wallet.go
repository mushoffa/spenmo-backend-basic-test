package entity

import (
	"time"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 03/11/2021
type Wallet struct {
	ID            string    `json:"id,omitempty"`
	Created       time.Time `json:"created,omitempty"`
	Updated       time.Time `json:"updated,omitempty"`
	Name          string    `json:"name,omitempty"`
	Balance       float64   `json:"balance,omitempty"`
	MaxLimit      float64   `json:"max_limit,omitempty"`
	AccountNumber string    `json:"account_number,omitempty"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (e *Wallet) ValidateBalance() error {
	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (e *Wallet) ValidateMaxLimit() error {
	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (e *Wallet) ValidateUserID() error {
	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (e *Wallet) Validate() error {
	return nil
}
