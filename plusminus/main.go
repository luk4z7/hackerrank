package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func plusMinus(arr []int32) {
	// Write your code here
	var positive, negative, zero []int32

	for i := range arr {
		if math.Signbit(float64(arr[i])) {
			negative = append(negative, arr[i])
		} else {
			if arr[i] == 0 {
				zero = append(zero, arr[i])
				continue
			}

			positive = append(positive, arr[i])
		}
	}

	fmt.Printf("%6f\n", float64(len(positive))/float64(len(arr)))
	fmt.Printf("%6f\n", float64(len(negative))/float64(len(arr)))
	fmt.Printf("%6f\n", float64(len(zero))/float64(len(arr)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
