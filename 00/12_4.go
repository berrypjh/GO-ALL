// 소인수분해
package main

import (
	"fmt"
)

func solution(n int) []int {
	result := []int{}

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			for {
				if n%i == 0 {
					n /= i
					continue
				}
				break
			}
			result = append(result, i)
		}
	}

	return result
}

func main() {
	result1 := solution(10000)
	fmt.Println(result1)

	result2 := solution(17)
	fmt.Println(result2)

	result3 := solution(420)
	fmt.Println(result3)
}
