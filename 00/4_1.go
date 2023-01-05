// 피자 나눠 먹기 (1)
package main

import "fmt"

func solution(n int) int {
	x := 1

	if n%7 == 0 {
		x = 0
	}

	return n/7 + x
}

func main() {
	result1 := solution(7)
	fmt.Println(result1)

	result2 := solution(1)
	fmt.Println(result2)

	result3 := solution(15)
	fmt.Println(result3)
}
