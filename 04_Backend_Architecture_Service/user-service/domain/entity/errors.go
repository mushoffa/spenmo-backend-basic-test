package entity

import (
	"errors"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 04/11/2021
var (
	ErrNameEmptyField = errors.New("Field name cannot be empty")

	ErrNameInvalidMinLength = errors.New("Field name must be at least 5 characters")

	ErrEmailEmptyField = errors.New("Field email cannot be empty")

	ErrEmailInvalid = errors.New("Invalid email format")

	ErrPhoneEmptyField = errors.New("Field phone number cannot be empty")

	ErrPhoneInvalidLength = errors.New("Field phone number must follow Indonesian mobile standard digit length")

	ErrPhoneInvalidPrefix = errors.New("Invalid mobile phone number prefix")

	ErrDOBEmptyField = errors.New("Field dob cannot be empty")

	ErrDOBInvalidFormat = errors.New("Field dob format must be correct DD/MM/YYYY")

	ErrUserIsExist = errors.New("User is already exists")
)