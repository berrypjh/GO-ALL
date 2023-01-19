// 삼각형의 완성조건 (2)
package main

import (
	"fmt"
)

func solution(sides []int) int {
	if sides[0] > sides[1] {
		sides[1], sides[0] = sides[0], sides[1]
	}

	return sides[0] + sides[0] - 1
}

func main() {
	result1 := solution([]int{1, 2})
	fmt.Println(result1)

	result2 := solution([]int{3, 6})
	fmt.Println(result2)

	result3 := solution([]int{11, 7})
	fmt.Println(result3)
}
