// 특이한 정렬
package main

import (
	"fmt"
	"math"
	"sort"
)

func solution(numlist []int, n int) []int {
	sort.Slice(numlist, func(i, j int) bool {
		x := math.Abs(float64(numlist[i] - n))
		y := math.Abs(float64(numlist[j] - n))

		if x == y {
			return numlist[i] > numlist[j]
		}

		return x < y
	})

	return numlist
}

func main() {
	result1 := solution([]int{1, 2, 3, 4, 5, 6}, 4)
	fmt.Println(result1)

	result2 := solution([]int{10000, 20, 36, 47, 40, 6, 10, 7000}, 30)
	fmt.Println(result2)
}
