// 겹치는 선분의 길이
package main

import (
	"fmt"
)

func solution(lines [][]int) int {
	m := map[int]int{}
	for _, v := range lines {
		for i := v[0]; i < v[1]; i++ {
			m[i]++
		}
	}

	result := 0
	for _, v := range m {
		if v > 1 {
			result++
		}
	}

	return result
}

func main() {
	result1 := solution([][]int{{0, 1}, {2, 5}, {3, 9}})
	fmt.Println(result1)

	result2 := solution([][]int{{-1, 1}, {1, 3}, {3, 9}})
	fmt.Println(result2)

	result3 := solution([][]int{{0, 5}, {3, 9}, {1, 10}})
	fmt.Println(result3)

	result4 := solution([][]int{{0, 2}, {-3, -1}, {-2, 1}})
	fmt.Println(result4)
}
