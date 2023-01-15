// 문자열 계산하기
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(my_string string) int {
	array := strings.Split(my_string, " ")
	num, _ := strconv.Atoi(array[0])
	result := num

	for i := 1; i < len(array); i += 2 {
		switch array[i] {
		case "+":
			n1, _ := strconv.Atoi(array[i+1])
			result += n1
		case "-":
			n2, _ := strconv.Atoi(array[i+1])
			result -= n2
		}
	}

	return result
}

func main() {
	result1 := solution("3 + 4")
	fmt.Println(result1)
}
