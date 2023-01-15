// 숫자 찾기
package main

import (
	"fmt"
	"strconv"
)

func solution(num int, k int) int {
	for i, v := range strconv.Itoa(num) {
		if string(v) == strconv.Itoa(k) {
			return i + 1
		}
	}

	return -1
}

func main() {
	result1 := solution(29183, 1)
	fmt.Println(result1)

	result2 := solution(232443, 4)
	fmt.Println(result2)

	result3 := solution(123456, 7)
	fmt.Println(result3)
}
