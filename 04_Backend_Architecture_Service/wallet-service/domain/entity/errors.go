package entity

import (
	"errors"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
var (
	ErrWalletIsExist = errors.New("Wallet is already exist")

	ErrInvalidUser = errors.New("Do not honor")

	ErrWalletMaxLimit = errors.New("Wallet max limit exceeded")
)
