package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	p = fmt.Println
)

func Solution1(S string) int {
	// write your code in Go 1.4
	split := strings.Split(S, ".")
	var max int
	for _, v := range split {
		spt := strings.Split(strings.Trim(v, " "), " ")
		var j int
		for _, x := range spt {
			if x == "" {
				continue
			}

			j++
		}

		if j > max {
			max = j
		}
	}

	return max
}

func Solution2(S string) string {
	dates := []string{}
	lines := []string{}
	result := make(map[string]string)
	city := make(map[string]int)

	data := strings.Split(S, "\\n")
	for i := range data {
		lines = append(lines, data[i])
		item := strings.Split(data[i], ",")
		dates = append(dates, strings.Trim(item[2], " "))

	}

	sort.Slice(dates, func(i, j int) bool {
		return dates[i] < dates[j]
	})

	for _, v := range dates {
		for i := range lines {
			if strings.Contains(lines[i], v) {
				item := strings.Split(data[i], ",")
				city[item[1]] += 1
				numb := strconv.Itoa(city[item[1]])

				name := item[1] + numb + item[0][len(item[0])-4:]
				result[lines[i]] = name
			}
		}
	}

	for i := range lines {
		p(result[lines[i]])
	}

	return ""
}

func main() {
	p(Solution1("We test coders. Give us a try?"))
	p(Solution1("Forget  CVs..Save time . x x"))

	data := `photo.jpg, Warsaw, 2013-09-05 14:08:15\njohn.png, London, 2015-06-20 15:13:22\nmyFriends.png, Warsaw, 2013-09-05 14:07:13\nEiffel.jpg, Paris, 2015-07-23 08:03:02\npisatower.jpg, Paris, 2015-07-22 23:59:59\nBOB.jpg, London, 2015-08-05 00:02:03\nnotredame.png, Paris, 2015-09-01 12:00:00\nme.jpg, Warsaw, 2013-09-06 15:40:22\na.png, Warsaw, 2016-02-13 13:33:50\nb.jpg, Warsaw, 2016-01-02 15:12:22\nc.jpg, Warsaw, 2016-01-02 14:34:30\nd.jpg, Warsaw, 2016-01-02 15:15:01\ne.png, Warsaw, 2016-01-02 09:49:09\nf.png, Warsaw, 2016-01-02 10:55:32\ng.jpg, Warsaw, 2016-02-29 22:13:11`

	p(Solution2(data))
}
