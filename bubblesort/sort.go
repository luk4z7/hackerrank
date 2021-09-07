package main

func bubbleSort(data []int32) []int32 {
	var swaps int64

	for i := range data {
		swap := false

		// len(data)-1 for get the next index
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swap = true
				swaps += 1
			}
		}

		if !swap {
			break
		}
	}

	p("swaps: ", swaps, "first: ", data[0], "last: ", data[len(data)-1])

	return data
}
