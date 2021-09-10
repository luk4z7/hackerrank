package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	p = fmt.Println
)

// Complete the minimumSwaps function below.
func minimumSwaps(data []int32) int32 {

	var arr []int
	for i := range data {
		arr = append(arr, int(data[i]))
	}

	var swaps int
	temp := map[int]int{}

	// make a backup of the arr
	for i, v := range arr {
		temp[v] = i
	}

	for i := range arr {
		// verify the differ between value of the arr[i] and the actual index
		if arr[i] != i+1 {
			swaps += 1

			t := arr[i]    // backup the value of the total that are in arr
			arr[i] = i + 1 // update the value of the arr with the number to actual index

			arr[temp[i+1]] = t  // get the value of the temp with actual index
			temp[t] = temp[i+1] // update the temp with the key: t and value of the temp
		}
	}

	fmt.Println(swaps)

	return int32(swaps)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
