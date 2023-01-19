// 숨어있는 숫자의 덧셈 (2)
package main

import (
	"fmt"
	"strconv"
)

func solution(my_string string) int {
	sum := 0

	for i := 0; i < len(my_string); {
		str := ""

		for i < len(my_string) {
			if '0' <= my_string[i] && my_string[i] <= '9' {
				str += string(my_string[i])
				i++
			} else {
				break
			}
		}

		if str != "" {
			num, _ := strconv.Atoi(str)
			sum += num
		}
		i++
	}

	return sum
}

func main() {
	// result1 := solution("aAb1B2cC34oOp")
	// fmt.Println(result1)

	// result2 := solution("1a2b3c4d123Z")
	// fmt.Println(result2)

	result3 := solution("231321313")
	fmt.Println(result3)
}
