// 피자 나눠 먹기 (2)
package main

import "fmt"

func solution(n int) int {
	x := 1
	for {
		if 6*x%n == 0 {
			return x
		}
		x++
	}
}

func main() {
	result1 := solution(6)
	fmt.Println(result1)

	result2 := solution(10)
	fmt.Println(result2)

	result3 := solution(4)
	fmt.Println(result3)
}
