package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getUsernames' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts INTEGER threshold as parameter.
 *
 * URL for cut and paste
 * https://jsonmock.hackerrank.com/api/article_users?page=<pageNumber>
 */

var (
	p = fmt.Println
)

const endpoint = "https://jsonmock.hackerrank.com/api/article_users?page="

type Response struct {
	Page       int64  `json:"page"`
	PerPage    int64  `json:"per_page"`
	Total      int64  `json:"total"`
	TotalPages int64  `json:"total_pages"`
	Data       []Data `json:"data"`
}

type Data struct {
	Id              int32  `json:"id"`
	UserName        string `json:"username"`
	About           string `json:"about"`
	Submitted       int32  `json:"submitted"`
	UpdatedAt       string `json:"updated_at"`
	SubmissionCount int32  `json:"submission_count"`
	CommentCount    int32  `json:"comment_count"`
	CreatedAt       int32  `json:"created_at"`
}

func getUsernames(threshold int32) []string {
	totalPages := 3
	data := []Data{}

	for i := 1; i <= totalPages; i++ {
		page := strconv.Itoa(i)
		response, err := http.Get(endpoint + page)
		if err != nil {
			panic(err)
		}

		respData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		object := Response{}
		json.Unmarshal(respData, &object)

		totalPages = int(object.TotalPages)
		data = append(data, object.Data...)
	}

	result := []string{}

	for i := range data {
		if data[i].SubmissionCount > threshold {
			p(data[i].UserName)

			result = append(result, data[i].UserName)
		}
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	thresholdTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	threshold := int32(thresholdTemp)

	result := getUsernames(threshold)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
