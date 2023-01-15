// n의 배수 고르기
package main

import (
	"fmt"
)

func solution(n int, numlist []int) []int {
	result := []int{}

	for _, v := range numlist {
		if v%n == 0 {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	result1 := solution(3, []int{4, 5, 6, 7, 8, 9, 10, 11, 12})
	fmt.Println(result1)

	result2 := solution(5, []int{1, 9, 3, 10, 13, 5})
	fmt.Println(result2)

	result3 := solution(12, []int{2, 100, 120, 600, 12, 12})
	fmt.Println(result3)
}
