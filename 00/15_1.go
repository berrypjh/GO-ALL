// 영어가 싫어요
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(numbers string) int64 {
	m := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for i, v := range m {
		numbers = strings.ReplaceAll(numbers, i, v)
	}

	result, _ := strconv.ParseInt(numbers, 10, 64)

	return result
}

func main() {
	result1 := solution("onetwothreefourfivesixseveneightnine")
	fmt.Println(result1)

	result2 := solution("onefourzerosixseven")
	fmt.Println(result2)
}
