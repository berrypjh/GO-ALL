// 제곱수 판별하기
package main

import (
	"fmt"
	"math"
)

func solution(n int) int {
	c := math.Sqrt(float64(n))

	if float64(int(c)) == math.Sqrt(float64(n)) {
		return 1
	}

	return 2
}

func main() {
	result1 := solution(144)
	fmt.Println(result1)

	result2 := solution(976)
	fmt.Println(result2)
}
