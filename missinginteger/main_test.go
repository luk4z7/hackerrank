package main

import (
	"fmt"
	"testing"
)

var (
	p = fmt.Println

	testCase = []struct {
		Data []int
		Want int
	}{
		{[]int{2, 3, 7, 6, 8, -1, -10, 15}, 1},
		{[]int{2, 3, -7, 6, 8, 1, -10, 15}, 4},
		{[]int{1, 1, 0, -1, -2}, 2},
	}
)

func TestSolution(t *testing.T) {
	for _, tc := range testCase {
		res := solution(tc.Data)
		if res != tc.Want {
			t.Errorf("Have = %d - Want = %d", res, tc.Want)
		}
	}
}

func TestFirstMissingPositive(t *testing.T) {
	for _, tc := range testCase {
		res := firstMissingPositive(tc.Data)
		if res != tc.Want {
			t.Errorf("Have = %d - Want = %d", res, tc.Want)
		}
	}
}
