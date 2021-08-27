package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	var max int64
	var min int64 = int64(arr[0])

	for _, v := range arr {
		v := int64(v)

		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	var sumMax, sumMin int64

	if max == min {
		sumMax = max * int64(len(arr)-1)
		sumMin = sumMax
	}

	for _, v := range arr {
		v := int64(v)

		if v != max {
			sumMax += v
		}

		if v != min {
			sumMin += v
		}
	}

	fmt.Println(sumMax, sumMin)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
