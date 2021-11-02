package entity

import (
	"errors"
)

var (
	ErrPANEmptyField = errors.New("Field PAN cannot be empty")

	ErrPANInvalidLength = errors.New("Field PAN must be exactly 16 characters")

	ErrCardIsExist = errors.New("Card is already exists")
)
