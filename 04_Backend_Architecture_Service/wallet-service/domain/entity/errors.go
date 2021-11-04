package entity

import (
	"errors"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 04/11/2021
var (
	ErrNameEmptyField = errors.New("Field name cannot be empty")

	ErrBalanceEmptyField = errors.New("Field balance cannot be empty")

	ErrMaxLimitEmptyField = errors.New("Field max_limit cannot be empty")

	ErrAccountNumberEmptyField = errors.New("Field account_number cannot be empty")

	ErrBalanceLessThanZero = errors.New("Field balance must be greater than zero")

	ErrMaxLimitLessThanZero = errors.New("Field max_limit must be greater than zero")

	ErrAccountNumberInvalidLength = errors.New("Field account_number must follow Indonesian mobile standard digit length")

	ErrAccountNumberInvalidPrefix = errors.New("Invalid mobile phone number prefix")

	ErrWalletIsExist = errors.New("Wallet is already exist")

	ErrInvalidUser = errors.New("Do not honor")

	ErrWalletMaxLimit = errors.New("Wallet max limit exceeded")
)
