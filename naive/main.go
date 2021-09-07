package main

import (
	"fmt"
)

var (
	p = fmt.Println
)

// https://www.geeksforgeeks.org/naive-algorithm-for-pattern-searching/
// https://github.com/GNaive/naive-rete-go
// https://github.com/crazyacking/algorithms-go
// https://github.com/amitkgupta/nearest_neighbour
//
// naive
// target 8
// lista 1 3 7 10
// out: 0 2

// target 462
// lista 9 8 20 9 10 442 13
// out 2 5

// target 30
// lista 9 8 20 9 10 442 13
// out 2 4
func main() {
	//list := []int{1, 3, 7, 10}
	list := []int{9, 8, 20, 9, 10, 442, 13}
	//target := 462
	//target := 8
	target := 30

	pos1 := naive1(list, target)
	p(pos1)

	pos2 := naive2(list, target)
	p(pos2)
}

// naive with only one for
func naive2(list []int, target int) []int {
	var positions []int
	var data []int

	for i := range list {
		data = append(data, list[i])

		if i != 0 {
			x := i - 1
			y := i + 1

			if len(list)-1 == i {
				y = i
			}

			if data[x]+list[y] == target {
				positions = append(positions, x)
				positions = append(positions, y)

				break
			}
		}
	}

	return positions
}

func naive1(list []int, target int) []int {
	var positions []int

	for i := 0; i <= len(list)-1; i++ {
		for j := 0; j < len(list); j++ {
			if list[i]+list[j] == target {

				positions = append(positions, i)
				break
			}
		}
	}

	return positions
}
