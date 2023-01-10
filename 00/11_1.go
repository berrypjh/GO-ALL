// 주사위의 개수
package main

import "fmt"

func solution(box []int, n int) int {
	result := 1

	for _, v := range box {
		result *= v / n
	}

	return result
}

func main() {
	result1 := solution([]int{1, 1, 1}, 1)
	fmt.Println(result1)

	result2 := solution([]int{10, 8, 6}, 3)
	fmt.Println(result2)
}
