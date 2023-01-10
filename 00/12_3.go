// 숨어있는 숫자의 덧셈 (1)
package main

import (
	"fmt"
	"strconv"
)

func solution(my_string string) int {
	result := 0

	for _, v := range my_string {
		if '0' <= v && v <= '9' {
			num, _ := strconv.Atoi(string(v))
			result += num
		}
	}

	return result
}

func main() {
	result1 := solution("aAb1B2cC34oOp")
	fmt.Println(result1)

	result2 := solution("1a2b3c4d123")
	fmt.Println(result2)
}
