package entity_test

import (
	"fmt"
	"testing"
	// "time"
	"user-service/domain/entity"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func TestUserName_Empty(t *testing.T) {
	user := entity.User{}
	err := user.ValidateName()
	if err != nil {
		fmt.Println(err)
	}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 04/11/2021
// @Updated
func TestDOB_Invalid(t *testing.T) {
	user := entity.User{DOB: "01-12-2000"}
	err := user.ValidateDateOfBirth()
	fmt.Println(err)
}