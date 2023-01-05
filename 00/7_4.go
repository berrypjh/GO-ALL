// 짝수의 합
package main

import "fmt"

func solution(n int) int {
	result := 0

	for i := 2; i <= n; i += 2 {
		result += i
	}

	return result
}

func main() {
	result1 := solution(10)
	fmt.Println(result1)

	result2 := solution(4)
	fmt.Println(result2)
}
