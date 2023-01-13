// 가까운 수
package main

import (
	"fmt"
	"math"
	"sort"
)

func solution(array []int, n int) int {
	result := 1
	temp := 100
	sort.Ints(array)

	for _, v := range array {
		x := math.Abs(float64(n - v))

		if int(x) < temp {
			temp = int(x)
			result = v
		} 
	}

	return result
}

func main() {
	result1 := solution([]int{3, 10, 28}, 20)
	fmt.Println(result1)

	result2 := solution([]int{10, 11, 12}, 13)
	fmt.Println(result2)
}
