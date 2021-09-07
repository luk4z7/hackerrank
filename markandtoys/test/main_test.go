package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestMaximumToys(t *testing.T) {
	c, err := ioutil.ReadFile("../../testcase.txt")
	if err != nil {
		panic(err)
	}

	pricesTemp := strings.Split(strings.TrimRight(string(c), "\r\n"), " ")

	var prices []int32
	n := 99383
	var k int32 = 806930886

	for i := 0; i < int(n); i++ {
		pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
		checkError(err)
		pricesItem := int32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}

	maximumToys(prices, k)
}
