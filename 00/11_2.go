// 합성수 찾기
package main

import "fmt"

func solution(n int) int {
	result := 0

	for i := 4; i <= n; i++ {
		for j := 2; j < i; j++ {
			if (i % j) == 0 {
				result++
				break
			}
		}
	}

	return result
}

func main() {
	result1 := solution(10)
	fmt.Println(result1)

	result2 := solution(15)
	fmt.Println(result2)

	result3 := solution(3)
	fmt.Println(result3)
}
