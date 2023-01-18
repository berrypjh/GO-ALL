// 최댓값 만들기 (2)
package main

import (
	"fmt"
)

func solution(numbers []int) int {
	result := -100000000

	for i, v := range numbers {
		for j := 0; j < i; j++ {
			if result < v*numbers[j] {
				result = v * numbers[j]
			}
		}
	}

	return result
}

func main() {
	result1 := solution([]int{1, 2, -3, 4, -5})
	fmt.Println(result1)

	result2 := solution([]int{0, -31, 24, 10, 1, 9})
	fmt.Println(result2)

	result3 := solution([]int{10, 20, 30, 5, 5, 20, 5})
	fmt.Println(result3)
}
