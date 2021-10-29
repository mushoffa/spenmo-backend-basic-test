package main

import (
	"fmt"
	
	"spenmo/usecase"
)

// @Created 29/10/2021
// @Updated
func main() {
	main := []int {20, 7, 8, 10, 2, 5, 6}

	fmt.Println(usecase.SequenceExists(main, []int{20, 7, 8, 10, 2, 5, 6}))
}