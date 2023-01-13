// 369게임
package main

import (
	"fmt"
	"strconv"
)

func solution(order int) int {
	str := strconv.Itoa(order)
	result := 0

	for _, v := range str {
		if v == '3' || v == '6' || v == '9' {
			result++
		}
	}

	return result
}

func main() {
	result1 := solution(3)
	fmt.Println(result1)

	result2 := solution(29423)
	fmt.Println(result2)
}
