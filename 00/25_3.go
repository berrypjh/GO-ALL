// 연속된 수의 합
package main

import (
	"fmt"
)

func solution(num int, total int) []int {
	result := []int{}
	n := (2*total/num - num + 1) / 2

	for i := 0; i < num; i++ {
		result = append(result, n+i)
	}

	return result
}

func main() {
	result1 := solution(3, 12)
	fmt.Println(result1)

	result2 := solution(5, 15)
	fmt.Println(result2)

	result3 := solution(4, 14)
	fmt.Println(result3)

	result4 := solution(5, 5)
	fmt.Println(result4)
}
