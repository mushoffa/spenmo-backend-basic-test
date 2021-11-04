package model

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 03/11/2021
type CreateWalletRequest struct {
	AccountNumber string  `json:"accountNumber" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	MaxLimit      float64 `json:"maxLimit" binding:"required"`
}
