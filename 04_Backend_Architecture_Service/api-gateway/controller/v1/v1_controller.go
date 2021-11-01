package v1

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type V1Controller struct {
	Card        CardController
	Transaction TransactionCont
	User        UserController
	Wallet      WalletController
}
