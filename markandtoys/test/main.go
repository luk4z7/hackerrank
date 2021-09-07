package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	//"sort"
	"strconv"
	"strings"
)

var (
	p = fmt.Println
)

/*
 * Complete the 'maximumToys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY prices
 *  2. INTEGER k
 */

func sweep(list []int32, swapped bool) ([]int32, bool) {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
			swapped = true
		}
	}

	return list, swapped
}

func sortInt(data []int32) []int32 {
	for i := 0; i < len(data); i++ {
		swapped := false
		data, swapped = sweep(data, swapped)

		if !swapped {
			break
		}
	}

	return data
}

func InsertionSortInt(list []int32) {
	var sorted []int32
	for _, item := range list {
		sorted = insert(sorted, item)
	}
	for i, val := range sorted {
		list[i] = val
	}
}

func insert(sorted []int32, item int32) []int32 {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			// https://github.com/golang/go/wiki/SliceTricks
			// 1, 2, 3      4    + 5, 6, 7
			// sorted[:i] + item + sorted[i:]
			// a = append(a[:i], append(make([]T, j), a[i:]...)...)
			return append(sorted[:i], append([]int32{item}, sorted[i:]...)...)
		}
	}

	return append(sorted, item)
}

func maximumToys(prices []int32, k int32) int32 {
	var total int32
	var value int64

	var sorted []int32
	for _, item := range prices {
		sorted = insert(sorted, item)
	}
	for i, val := range sorted {
		prices[i] = val
	}

	for _, v := range prices {
		if v < k {
			if value+int64(v) > int64(k) {
				continue
			}

			value += int64(v)
			total += 1
		}
	}

	// var newPrices []int

	// for _, v := range prices {
	// 	newPrices = append(newPrices, int(v))
	// }

	// sort.Ints(newPrices)
	// prices = sortInt(prices)

	//	for sweep := 0; sweep < len(prices); sweep++ {
	//		swapped := false
	//		for i := 0; i < len(prices)-1-sweep; i++ {
	//			if prices[i] > prices[i+1] {
	//				prices[i], prices[i+1] = prices[i+1], prices[i]
	//				swapped = true
	//			}
	//		}
	//
	//		if !swapped {
	//			break
	//		}
	//	}

	p(total)

	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	c, err := ioutil.ReadFile("testcase.txt")
	if err != nil {
		panic(err)
	}

	pricesTemp := strings.Split(strings.TrimRight(string(c), "\r\n"), " ")

	var prices []int32

	for i := 0; i < int(n); i++ {
		pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
		checkError(err)
		pricesItem := int32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}

	result := maximumToys(prices, k)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
