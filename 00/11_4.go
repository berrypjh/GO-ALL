// 팩토리얼
package main

import "fmt"

func solution(n int) int {
	result := 0
	num := 1

	for i := 1; i <= 10; i++ {
		num *= i
		if num <= n {
			result = i
		}
	}

	return result
}

func main() {
	result1 := solution(3628800)
	fmt.Println(result1)

	result2 := solution(7)
	fmt.Println(result2)
}
