// 순서쌍의 개수
package main

import (
	"fmt"
)

func solution(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result++
		}
	}

	return result
}

func main() {
	result1 := solution(20)
	fmt.Println(result1)

	result2 := solution(100)
	fmt.Println(result2)
}
