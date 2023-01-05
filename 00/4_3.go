// 피자 나눠 먹기 (3)
package main

import "fmt"

func solution(slice int, n int) int {
	result := n / slice

	if n%slice != 0 {
		result++
	}

	return result
}

func main() {
	result1 := solution(7, 10)
	fmt.Println(result1)

	result2 := solution(4, 12)
	fmt.Println(result2)
}
