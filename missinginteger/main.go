package main

// https://www.geeksforgeeks.org/find-the-smallest-positive-number-missing-from-an-unsorted-array/

func max(d []int) int {
	var max int
	for i := range d {
		if i > max {
			max = i
		}
	}

	return max
}

func inArray(v int, d []int) bool {
	for _, k := range d {
		if k == v {
			return true
		}
	}

	return false
}

func solution(a []int) int {
	if max(a) <= 0 {
		return 1
	}

	i := -1000001
	c := true

	for c {
		i++
		if (i > 0 && !inArray(i, a)) || i > 1000000 {
			c = false
		}
	}

	if i <= 1000000 {
		if i < 1 {
			return 1
		}

		return i
	}

	return 0
}

func firstMissingPositive(a []int) int {
	total := len(a)

	// Loop to traverse the whole array
	for i := 0; i < total; i++ {
		// Loop to check boundary
		// condition and for swapping
		for a[i] >= 1 && a[i] <= total && a[i] != a[a[i]-1] {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}

	// Checking any element which
	// is not equal to i+1
	for i := 0; i < total; i++ {
		if a[i] != i+1 {
			return i + 1
		}
	}

	// Nothing is present return last index
	return total + 1
}
