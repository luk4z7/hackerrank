package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "sort"
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

func maximumToys(prices []int32, k int32) int32 {
	var total int32
	var value int64

	for sweep := 0; sweep < len(prices); sweep++ {
		swapped := false
		for i := 0; i < len(prices)-1-sweep; i++ {
			if prices[i] > prices[i+1] {
				prices[i], prices[i+1] = prices[i+1], prices[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
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

	pricesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

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
