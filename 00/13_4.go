// 삼각형의 완성조건 (1)
package main

import (
	"fmt"
	"sort"
)

func solution(sides []int) int {
	sort.Ints(sides)

	if sides[2] < sides[1]+sides[0] {
		return 1
	}

	return 2
}

func main() {
	result1 := solution([]int{1, 2, 3})
	fmt.Println(result1)

	result2 := solution([]int{3, 6, 2})
	fmt.Println(result2)

	result3 := solution([]int{199, 72, 222})
	fmt.Println(result3)
}
