// 등수 매기기
package main

import (
	"fmt"
	"sort"
)

func solution(score [][]int) []int {
	result := make([]int, len(score))
	temp := []float64{}

	for _, v := range score {
		n := float64(v[0]+v[1]) / 2
		temp = append(temp, n)
	}

	sort.Float64s(temp)

	for i, v := range score {
		n := float64(v[0]+v[1]) / 2
		for j, w := range temp {
			if n == w {
				result[i] = len(temp) - j
			}
		}
	}

	return result
}

func main() {
	result1 := solution([][]int{{80, 70}, {90, 50}, {40, 70}, {50, 80}})
	fmt.Println(result1)

	result2 := solution([][]int{{80, 70}, {70, 80}, {30, 50}, {90, 100}, {100, 90}, {100, 100}, {10, 30}})
	fmt.Println(result2)

	result3 := solution([][]int{{1, 3}, {3, 1}, {2, 3}, {3, 2}, {1, 0}})
	fmt.Println(result3)
}
