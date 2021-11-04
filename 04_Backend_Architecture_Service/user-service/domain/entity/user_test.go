package entity_test

import (
	"fmt"
	"testing"
	"time"
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

func TestTime(t *testing.T) {
	today := time.Now()
	dateLayout := "02-01-2006"
	timeLayout := "15:04:05"
	
	fmt.Println("Date: ", today.Format(dateLayout))
	fmt.Println("Time: ", today.Format(timeLayout))
}
