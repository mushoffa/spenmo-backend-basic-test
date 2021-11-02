package entity

import (
	"time"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type Card struct {
	ID           string    `json:"id"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
	PAN          string    `json:"pan"`
	Name         string    `json:"name"`
	ExpiryDate   string    `json:"expiry_date"`
	LimitDaily   float64   `json:"limit_daily"`
	LimitMonthly float64   `json:"limit_monthly"`
	CardType     string    `json:"cardType"`
	Status       string    `json:"status"`
	PIN          string    `json:"pin"`
	WalletID     string    `json:"wallet_id"`
	UserID       string    `json:"user_id"`
	// PIN 			uint32 		`json:"pin"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (c *Card) ValidatePAN() error {
	if c.PAN == "" {
		return ErrPANEmptyField
	}

	if len(c.PAN) != 16 {
		return ErrPANInvalidLength
	}

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (c *Card) Validate() error {
	if err := c.ValidatePAN(); err != nil {
		return err
	}

	return nil
}
