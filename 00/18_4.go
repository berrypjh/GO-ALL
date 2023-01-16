// 문자열 정렬하기 (2)
package main

import (
	"fmt"
	"sort"
	"strings"
)

func solution(my_string string) string {
	str := strings.ToLower(my_string)
	array := []int{}

	for _, v := range str {
		array = append(array, int(v))
	}

	result := ""
	sort.Ints(array)
	for _, v := range array {
		result += string(v)
	}

	return result
}

func main() {
	result1 := solution("Bcad")
	fmt.Println(result1)

	result2 := solution("heLLo")
	fmt.Println(result2)

	result3 := solution("Python")
	fmt.Println(result3)
}
