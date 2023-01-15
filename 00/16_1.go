// 편지
package main

import (
	"fmt"
)

func solution(message string) int {
	return len(message) * 2
}

func main() {
	result1 := solution("happy birthday!")
	fmt.Println(result1)

	result2 := solution("I love you~")
	fmt.Println(result2)
}
