package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	hour := s[:2]           // 07
	time := s[2 : len(s)-2] // :05:45
	h := s[len(s)-2:]       // PM/AM
	d := s[:len(s)-2]       // default

	pm := map[string]string{
		"01": "13",
		"02": "14",
		"03": "15",
		"04": "16",
		"05": "17",
		"06": "18",
		"07": "19",
		"08": "20",
		"09": "21",
		"10": "22",
		"11": "23",
		"12": "12",
	}

	am := map[string]string{
		"12": "00",
	}

	switch h {
	case "PM":
		s = pm[hour] + time
	case "AM":
		s = d

		if am[hour] != "" {
			s = am[hour] + time
		}
	}

	fmt.Println(s)

	return s
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	stdout, err := os.Create("./teste")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
