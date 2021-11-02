package v1

import (
	"github.com/mushoffa/spenmo-proto/protos"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type V1Controller struct {
	Card        CardController
	Transaction TransactionController
	User        UserController
	Wallet      WalletController
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func NewV1Controller(user protos.UserServiceClient, wallet protos.WalletServiceClient, card protos.CardServiceClient) *V1Controller {
	return &V1Controller{
		User:   NewUserController(user),
		Wallet: NewWalletController(wallet),
		Card:   NewCardController(card),
		Transaction: NewTransactionController(),
	}
}
