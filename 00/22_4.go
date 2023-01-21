// 유한소수 판별하기
package main

import (
	"fmt"
)

func solution(a int, b int) int {
	div := 1

	for i := 1; i < b; i++ {
		if a%i == 0 && b%i == 0 {
			div = i
		}
	}

	x := b / div
	for {
		if x%2 == 0 {
			x /= 2
		} else if x%5 == 0 {
			x /= 5
		} else {
			if x == 1 {
				return 1
			}
			return 2
		}
	}
}

func main() {
	result1 := solution(7, 20)
	fmt.Println(result1)

	result2 := solution(11, 22)
	fmt.Println(result2)

	result3 := solution(12, 21)
	fmt.Println(result3)
}
