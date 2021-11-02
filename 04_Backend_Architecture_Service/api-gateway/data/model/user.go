package model

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type RegisterUserRequest struct {
	Name 		string 	`json:"name" binding:"required"`
	PhoneNumber string 	`json:"phoneNumber" binding:"required"`
	Email 		string 	`json:"email" binding:"required"`
	DOB 		string 	`json:"dob" binding:"required"`
}