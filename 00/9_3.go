// 구슬을 나누는 경우의 수
package main

import "fmt"

func solution(balls int, share int) int {
	result := 1

	for i := 0; i < share; i++ {
		result *= balls - i
		result /= i + 1
	}

	return result
}

func main() {
	result1 := solution(3, 2)
	fmt.Println(result1)

	result2 := solution(5, 3)
	fmt.Println(result2)

	result3 := solution(30, 2)
	fmt.Println(result3)

	result4 := solution(15, 10)
	fmt.Println(result4)
}
