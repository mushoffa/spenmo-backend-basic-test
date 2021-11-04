package entity

import (
	"time"
	"net/mail"
	"regexp"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 02/11/2021
type User struct {
	ID          string    `json:"id"`
	Created     time.Time `json:"created"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	DOB         string    `json:"dob"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (u *User) ValidateName() error {

	if u.Name == "" {
		return ErrNameEmptyField
	}

	if len(u.Name) < 5 {
		return ErrNameInvalidMinLength
	}

	if len(u.Name) >= 25 {
		u.Name = u.Name[0:25]
	}

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated 04/11/2021
func (u *User) ValidateEmail() error {
	if u.Email == "" {
		return ErrEmailEmptyField
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return ErrEmailInvalid
	}

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 04/11/2021
// @Updated 
func (u *User) ValidatePhoneNumber() error {
	if u.PhoneNumber == "" {
		return ErrPhoneEmptyField
	}	

	digit := len(u.PhoneNumber)
	if (digit < 10) || (digit > 12) {
		return ErrPhoneInvalidLength
	}

	// Partially correct, better using regex
	if u.PhoneNumber[0:2] != "08" {
		return ErrPhoneInvalidPrefix
	}

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 04/11/2021
// @Updated 
func (u *User) ValidateDateOfBirth() error {
	if u.DOB == "" {
		return ErrDOBEmptyField
	}

	regex := regexp.MustCompile("^(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d)")
 
 	isValid := regex.MatchString(u.DOB)
 	
	if !isValid {
		return ErrDOBInvalidFormat
	} 

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (u *User) Validate() error {

	if err := u.ValidateName(); err != nil {
		return err
	}

	if err := u.ValidateEmail(); err != nil {
		return err
	}

	if err := u.ValidatePhoneNumber(); err != nil {
		return err
	}

	if err := u.ValidateDateOfBirth(); err != nil {
		return err
	}

	return nil
}
