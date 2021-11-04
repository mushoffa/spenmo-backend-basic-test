package model

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
type CreateBillRequest struct {
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type PurchaseRequest struct {
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 03/11/2021
type TopupRequest struct {
	AccountNumber string  `json:"accountNumber" binding:"required"`
	WalletID      string  `json:"walletId" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
}
