package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"A", "C", "C"}

	fmt.Println(sort.SearchStrings(a, "A")) // 0
	fmt.Println(sort.SearchStrings(a, "B")) // 1
	fmt.Println(sort.SearchStrings(a, "C")) // 1
	fmt.Println(sort.SearchStrings(a, "D")) // 3
}
