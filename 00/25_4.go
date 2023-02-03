// 다음에 올 숫자
package main

import (
	"fmt"
)

func solution(common []int) int {
	if common[1]-common[0] == common[2]-common[1] {
		return common[len(common)-1] + common[1] - common[0]
	}

	return common[len(common)-1] * common[1] / common[0]
}

func main() {
	result1 := solution([]int{1, 2, 3, 4})
	fmt.Println(result1)

	result2 := solution([]int{2, 4, 8})
	fmt.Println(result2)
}
