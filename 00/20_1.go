// 직사각형 넓이 구하기
package main

import (
	"fmt"
	"math"
)

func solution(dots [][]int) int {
	x1, y1 := dots[0][0], dots[0][1]
	x2, y2 := 0, 0

	for i := 0; i < len(dots); i++ {
		if x1 != dots[i][0] {
			x2 = dots[i][0]
		}

		if y1 != dots[i][1] {
			y2 = dots[i][1]
		}
	}

	n1 := math.Abs(float64(x1 - x2))
	n2 := math.Abs(float64(y1 - y2))

	return int(n1 * n2)
}

func main() {
	result1 := solution([][]int{[]int{1, 2}, []int{2, 1}, []int{2, 2}, []int{1, 2}})
	fmt.Println(result1)

	result2 := solution([][]int{[]int{-1, -1}, []int{1, 1}, []int{1, -1}, []int{-1, 1}})
	fmt.Println(result2)
}
