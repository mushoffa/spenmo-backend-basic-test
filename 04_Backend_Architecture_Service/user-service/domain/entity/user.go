package entity

import (
	"time"
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
// @Updated
func (u *User) Validate() error {

	if err := u.ValidateName(); err != nil {
		return err
	}

	return nil
}
