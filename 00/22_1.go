// 저주의 숫자 3
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		result++
		for result%3 == 0 || strings.Contains(strconv.Itoa(result), "3") {
			fmt.Println("숫자 ", result)
			result++
		}
	}

	return result
}

func main() {
	result1 := solution(15)
	fmt.Println(result1)

	result2 := solution(40)
	fmt.Println(result2)
}
