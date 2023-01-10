// 문자열 정렬하기 (1)
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func solution(my_string string) []int {
	result := []int{}

	for _, v := range my_string {
		if 48 <= v && v <= 57 {
			num, err := strconv.Atoi(string(v))

			if err != nil {
				fmt.Println(err)
			}

			result = append(result, num)
		}
	}

	sort.Ints(result)
	return result
}

func main() {
	result1 := solution("hi12392")
	fmt.Println(result1)

	result2 := solution("p2o4i8gj2")
	fmt.Println(result2)

	result3 := solution("abcde0")
	fmt.Println(result3)
}
