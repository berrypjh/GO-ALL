// 겹치는 선분의 길이
package main

import (
	"fmt"
)

func solution(lines [][]int) int {
	return 0
}

func main() {
	result1 := solution([][]int{[]int{0, 1}, []int{2, 5}, []int{3, 9}})
	fmt.Println(result1)

	result2 := solution([][]int{[]int{-1, 1}, []int{1, 3}, []int{3, 9}})
	fmt.Println(result2)

	result3 := solution([][]int{[]int{0, 5}, []int{3, 9}, []int{1, 10}})
	fmt.Println(result3)
}
