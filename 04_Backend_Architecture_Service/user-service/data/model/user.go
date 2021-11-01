package model

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type UserDB struct {
	ID          string `gorm:"id"`
	Name        string `gorm:"name"`
	PhoneNumber string `gorm:"phone_number"`
	Email       string `gorm:"email"`
	DOB         string `gorm:"dob"`
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func (t *UserDB) TableName() string {
	return "users"
}
