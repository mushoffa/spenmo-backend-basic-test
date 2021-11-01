package entity

import (
	"errors"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
var (
	ErrNameEmptyField = errors.New("Field name cannot be empty")

	ErrNameInvalidMinLength = errors.New("Field name must be at least 5 characters")

	ErrUserIsExist = errors.New("User is already exists")
)