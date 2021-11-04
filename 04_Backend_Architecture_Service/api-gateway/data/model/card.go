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
// @Updated 04/11/2021
type LinkCardRequest struct {
	PAN 			string 	`json:"pan" binding:"required"`
	WalletID 		string 	`json:"walletId" binding:"required"`
	AccountNumber 	string 	`json:"accountNumber" binding:"required"`
}