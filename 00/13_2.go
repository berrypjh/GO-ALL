// 배열 원소의 길이
package main

import (
	"fmt"
)

func solution(strlist []string) []int {
	result := []int{}

	for _, v := range strlist {
		result = append(result, len(v))
	}

	return result
}

func main() {
	result1 := solution([]string{"We", "are", "the", "world!"})
	fmt.Println(result1)

	result2 := solution([]string{"I", "Love", "Programmers."})
	fmt.Println(result2)
}
