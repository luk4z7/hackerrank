package main

import (
	"fmt"
	"sort"
)

var (
	p = fmt.Println
)

// https://yourbasic.org/golang/find-search-contains-slice/
// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}

func Find(a []int, x int) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}

	return len(a)
}

func main() {
	a := []int{1, 2, 3, 4, 6, 7, 8}

	p(Find(a, 6))

	// binary search
	p(sort.SearchInts(a, 6))
}
