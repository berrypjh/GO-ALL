// 세균 증식
package main

import (
	"fmt"
)

func solution(n int, t int) int {
	result := n

	for i := 1; i <= t; i++ {
		result *= 2
	}

	return result
}

func main() {
	result1 := solution(2, 10)
	fmt.Println(result1)

	result2 := solution(7, 15)
	fmt.Println(result2)
}
