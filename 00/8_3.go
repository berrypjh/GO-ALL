// 진료 순서 정하기
package main

import (
	"fmt"
)

func solution(emergency []int) []int {
	result := make([]int, len(emergency))

	for i := range emergency {
		for j := range emergency {
			if emergency[i] < emergency[j] {
				result[i]++
			}
		}
		result[i]++
	}

	return result
}

func main() {
	result1 := solution([]int{3, 76, 24})
	fmt.Println(result1)

	result2 := solution([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(result2)

	result3 := solution([]int{30, 10, 23, 6, 100})
	fmt.Println(result3)
}
