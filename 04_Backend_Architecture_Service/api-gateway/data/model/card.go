package model

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type CreateCardRequest struct {
	PAN 		 string 	`json:"pan" binding:"required"`
	ExpiryDate 	 string 	`json:"expiryDate" binding:"required"`
	CardType 	 string 	`json:"cardType" binding:"required"`
	LimitDaily   float64 	`json:"limitDaily" binding:"required"`
	LimitMonthly float64 	`json:"limitMonthly" binding:"required"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type LinkCardRequest struct {
	CardID 		string 	`json:"cardId" binding:"required"`
	WalletID 	string 	`json:"walletId" binding:"required"`
	UserID 		string 	`json:"userId" binding:"required"`
}