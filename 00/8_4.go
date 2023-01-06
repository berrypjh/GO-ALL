// 순서쌍의 개수
package main

import (
	"fmt"
)

func solution(emergency []int) []int {
	result := []int{}

	for i := 0; i < len(emergency); i++ {
		if len(emergency)%i == 0 {
			result = append(result, i)
		}
	}

	num := len(result)
	return num
}

func main() {
	result1 := solution([]int{3, 76, 24})
	fmt.Println(result1)

	result2 := solution([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(result2)

	result3 := solution([]int{30, 10, 23, 6, 100})
	fmt.Println(result3)
}
