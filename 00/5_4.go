// 배열 뒤집기
package main

import "fmt"

func solution(num_list []int) []int {
	result := []int{}

	for i := range num_list {
		result = append(result, num_list[len(num_list)-i-1])
	}

	return result
}

func main() {
	result1 := solution([]int{1, 2, 3, 4, 5})
	fmt.Println(result1)

	result2 := solution([]int{1, 1, 1, 1, 1, 2})
	fmt.Println(result2)

	result3 := solution([]int{1, 0, 1, 1, 1, 3, 5})
	fmt.Println(result3)
}
