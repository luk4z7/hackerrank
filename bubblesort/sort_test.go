package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

var (
	p = fmt.Println
)

func TestBubbleSort(t *testing.T) {
	c, err := ioutil.ReadFile("../testcase.txt")
	if err != nil {
		panic(err)
	}

	dataTemp := strings.Split(strings.TrimRight(string(c), "\r\n"), " ")

	var data []int32
	n := 99383

	for i := 0; i < int(n); i++ {
		dataItemTemp, err := strconv.ParseInt(dataTemp[i], 10, 64)
		if err != nil {
			panic(err)
		}

		dataItem := int32(dataItemTemp)
		data = append(data, dataItem)
	}

	bubbleSort(data)
}
