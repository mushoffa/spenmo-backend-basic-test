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
// @Updated 04/11/2021
func (e *Wallet) ValidateName() error {
	if e.Name == "" {
		return ErrNameEmptyField
	}

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated 04/11/2021
func (e *Wallet) ValidateBalance() error {
	// if e.Balance == "" {
	// 	return ErrBalanceEmptyField
	// }

	if e.Balance < 0 {
		return ErrBalanceLessThanZero
	}

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated 04/11/2021
func (e *Wallet) ValidateMaxLimit() error {
	// if e.MaxLimit == "" {
	// 	return ErrMaxLimitEmptyField
	// }

	if e.MaxLimit < 0 {
		return ErrMaxLimitLessThanZero
	}

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated 04/11/2021
func (e *Wallet) ValidateAccountNumber() error {
	if e.AccountNumber == "" {
		return ErrAccountNumberEmptyField
	}

	digit := len(e.AccountNumber)
	if (digit < 10) || (digit > 12) {
		return ErrAccountNumberInvalidLength
	}

	// Partially correct, better using regex
	if e.AccountNumber[0:2] != "08" {
		return ErrAccountNumberInvalidPrefix
	}
	
	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated 04/11/2021
func (e *Wallet) Validate() error {
	if err := e.ValidateName(); err != nil {
		return err
	}

	if err := e.ValidateBalance(); err != nil {
		return err
	}

	if err := e.ValidateMaxLimit(); err != nil {
		return err
	}

	if err := e.ValidateAccountNumber(); err != nil {
		return err
	}
	return nil
}
