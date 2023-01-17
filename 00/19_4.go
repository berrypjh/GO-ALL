// 머쓱이보다 키 큰 사람
package main

import (
	"fmt"
)

func solution(array []int, height int) int {
	result := 0

	for _, v := range array {
		if v > height {
			result++
		}
	}

	return result
}

func main() {
	result1 := solution([]int{149, 180, 192, 170}, 167)
	fmt.Println(result1)

	result2 := solution([]int{180, 120, 140}, 190)
	fmt.Println(result2)
}
