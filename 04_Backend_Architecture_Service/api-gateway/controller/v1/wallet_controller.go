package v1

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type WalletController interface {
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type walletController struct {
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewWalletController() WalletController {
	return &walletController{}
}
