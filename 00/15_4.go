// 약수 구하기
package main

import (
	"fmt"
)

func solution(n int) []int {
	result := []int{}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	result1 := solution(24)
	fmt.Println(result1)

	result2 := solution(29)
	fmt.Println(result2)
}
