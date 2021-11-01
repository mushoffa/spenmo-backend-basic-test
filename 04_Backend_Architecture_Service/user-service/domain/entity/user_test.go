package entity_test

import (
	"fmt"
	"testing"
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
