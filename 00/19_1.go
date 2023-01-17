// 7의 개수
package main

import (
	"fmt"
	"strconv"
)

func solution(array []int) int {
	result := 0

	for _, v := range array {
		s := strconv.Itoa(v)

		for _, x := range s {
			if x == '7' {
				result++
			}
		}
	}

	return result
}

func main() {
	result1 := solution([]int{7, 77, 17})
	fmt.Println(result1)

	result2 := solution([]int{10, 29})
	fmt.Println(result2)
}
