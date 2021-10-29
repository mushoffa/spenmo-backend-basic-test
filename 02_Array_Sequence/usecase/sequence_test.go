package usecase_test

import (
	"fmt"
	"testing"

	"spenmo/usecase"

	"github.com/stretchr/testify/assert"
)

var dataSequence = []int{20, 7, 8, 10, 2, 5, 6}

// @Created 29/10/2021
// @Updated
func TestSequenceExists_Success(t *testing.T) {
	testSuccess := []struct {
		sequence []int
	}{
		{
			sequence: []int{20,7},
		},
		{
			sequence: []int{7,8},
		},
		{
			sequence: []int{8,10},
		},
		{
			sequence: []int{10,2},
		},
		{
			sequence: []int{2,5},
		},
		{
			sequence: []int{5,6},
		},
		{
			sequence: []int{20, 7, 8, 10, 2, 5, 6},
		},
	}

	for i, _ := range testSuccess {
		tc := testSuccess[i]
		t.Run(fmt.Sprintf("%v",tc.sequence), func(t *testing.T) {
			t.Parallel()
			result := usecase.SequenceExists(dataSequence, tc.sequence)
			assert.Equal(t, result, true)
		})
	}
}

// @Created 29/10/2021
// @Updated
func TestSequenceExists_Failed(t *testing.T) {
	testFailed := []struct {
		sequence []int
	}{
		{
			sequence: []int{8,7},
		},
		{
			sequence: []int{7,10},
		},
		{
			sequence: []int{5,20},
		},
		{
			sequence: []int{10,5},
		},
		{
			sequence: []int{2,10},
		},
		{
			sequence: []int{20,6},
		},
		{
			sequence: []int{20, 7, 8, 10, 2, 6, 5},
		},
		{
			sequence: []int{20, 7, 8, 10, 2, 6, 5, 5, 4},
		},
		{
			sequence: []int{91,92,93,94,95,96,97},
		},
	}

	for i, _ := range testFailed {
		tc := testFailed[i]
		t.Run(fmt.Sprintf("%v",tc.sequence), func(t *testing.T) {
			t.Parallel()
			result := usecase.SequenceExists(dataSequence, tc.sequence)
			assert.Equal(t, result, false)
		})
	}
}