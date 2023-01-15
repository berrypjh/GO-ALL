// n의 배수 고르기
package main

import (
	"fmt"
	"strconv"
)

func solution(n int) int {
	result := 0

	for _, v := range strconv.Itoa(n) {
		i, _ := strconv.Atoi(string(v))
		result += i
	}

	return result
}

func main() {
	result1 := solution(1234)
	fmt.Println(result1)

	result2 := solution(930211)
	fmt.Println(result2)
}
