package v1

import "github.com/mushoffa/spenmo-proto/protos"

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type WalletController interface {
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type walletController struct {
	client protos.WalletServiceClient
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewWalletController(client protos.WalletServiceClient) WalletController {
	return &walletController{client}
}
