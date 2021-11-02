package model

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type CreateWalletRequest struct {
	UserID 		string 		`json:"userId" binding:"required"`
	Name 	 	string 		`json:"name" binding:"required"`
	MaxLimit 	float64 	`json:"maxLimit" binding:"required"`
}