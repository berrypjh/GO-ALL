// 양꼬치
package main

import "fmt"

func solution(n int, k int) int {
	sum := n*12000 + k*2000

	if n >= 10 {
		sum -= n / 10 * 2000
	}

	return sum
}

func main() {
	result1 := solution(10, 3)
	fmt.Println(result1)

	result2 := solution(64, 6)
	fmt.Println(result2)
}
