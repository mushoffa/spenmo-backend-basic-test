package entity

import (
	"time"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type Card struct {
	ID 				string 		`json:"id"`
	Created 		time.Time 	`json:"created"`
	Updated 		time.Time 	`json:"updated"`
	PAN 			string 		`json:"pan"`
	Name 			string 		`json:"name"`
	LimitDaily 		float64 	`json:"limitDaily"`
	LimitMonthly 	float64 	`json:"limitMonthly"`
	CardType 		string 		`json:"cardType"`
	Status 			string 		`json:"status"`
	// PIN 			uint32 		`json:"pin"`
}