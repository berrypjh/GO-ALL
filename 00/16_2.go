// 가장 큰 수 찾기
package main

import (
	"fmt"
)

func solution(array []int) []int {
	result := []int{0, 0}

	for i := 0; i < len(array); i++ {
		if result[0] < array[i] {
			result[0] = array[i]
			result[1] = i
		}
	}
	return result
}

func main() {
	result1 := solution([]int{1, 8, 3})
	fmt.Println(result1)

	result2 := solution([]int{9, 10, 11, 8})
	fmt.Println(result2)
}
