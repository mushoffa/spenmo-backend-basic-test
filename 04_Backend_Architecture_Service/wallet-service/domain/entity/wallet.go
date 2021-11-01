package entity

import (
	"time"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type Wallet struct {
	ID string `json:"id,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Name string `json:"name,omitempty"`
	Balance float64 `json:"balance,omitempty"`
	UserID string `json:"userId,omitempty"`
}