// 2차원으로 만들기
package main

import "fmt"

func solution(num_list []int, n int) [][]int {
	result := make([][]int, len(num_list)/n)

	for i := 0; i < len(num_list)/n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = num_list[j+(n*i)]
		}
	}

	return result
}

func main() {
	result1 := solution([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2)
	fmt.Println(result1)

	result2 := solution([]int{100, 95, 2, 4, 5, 6, 18, 33, 948}, 3)
	fmt.Println(result2)
}
