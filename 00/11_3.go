// 최댓값 만들기 (1)
package main

import "fmt"

func solution(numbers []int) int {
	result := 0

	for i, v := range numbers {
		for j := 0; j < i; j++ {
			if result < v * numbers[j] {
				result = v * numbers[j]
			}
		}
	}

    return result
}

func main() {
	result1 := solution([]int{1, 2, 3, 4, 5})
	fmt.Println(result1)

	result2 := solution([]int{0, 31, 24, 10, 1, 9})
	fmt.Println(result2)
}
