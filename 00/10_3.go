// 공 던지기
package main

import "fmt"

func solution(numbers []int, k int) int {
	return numbers[(k - 1) * 2 % len(numbers)]
	// return ((k - 1) * 2 % len(numbers)) + 1
}

func main() {
	result1 := solution([]int{1, 2, 3, 4}, 2)
	fmt.Println(result1)

	result2 := solution([]int{1, 2, 3, 4, 5, 6}, 5)
	fmt.Println(result2)

	result3 := solution([]int{1, 2, 3}, 3)
	fmt.Println(result3)
}
