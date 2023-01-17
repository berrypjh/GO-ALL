// 중복된 숫자 개수
package main

import (
	"fmt"
)

func solution(array []int, n int) int {
	result := 0

	for _, v := range array {
		if v == n {
			result++
		}
	}

	return result
}

func main() {
	result1 := solution([]int{1, 1, 2, 3, 4, 5}, 1)
	fmt.Println(result1)

	result2 := solution([]int{0, 2, 3, 4}, 1)
	fmt.Println(result2)
}
