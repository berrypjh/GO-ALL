// 평행
package main

import (
	"fmt"
)

func solution(dots [][]int) int {
	temp := [][]int{
		{0, 1, 2, 3},
		{0, 2, 3, 1},
		{0, 3, 1, 2},
	}

	for _, v := range temp {
		m1 := float64(dots[v[0]][1]-dots[v[1]][1]) / float64(dots[v[0]][0]-dots[v[1]][0])
		m2 := float64(dots[v[2]][1]-dots[v[3]][1]) / float64(dots[v[2]][0]-dots[v[3]][0])

		if m1 == m2 {
			return 1
		}
	}

	return 0
}

func main() {
	result1 := solution([][]int{{1, 4}, {9, 2}, {3, 8}, {11, 6}})
	fmt.Println(result1)

	result2 := solution([][]int{{3, 5}, {4, 1}, {2, 4}, {5, 10}})
	fmt.Println(result2)
}
