package entity

import "errors"

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	DOB         string `json:"dob"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (u *User) ValidateName() error {

	if u.Name == "" {
		return errors.New("Field name cannot be empty")
	}

	if len(u.Name) < 5 {
		return errors.New("Field name must be at least 5 characters")
	}
	return nil
}
