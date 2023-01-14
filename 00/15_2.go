// 인덱스 바꾸기
package main

import (
	"fmt"
)

func solution(my_string string, num1 int, num2 int) string {
	result := []rune(my_string)
	result[num1], result[num2] = result[num2], result[num1]

	return string(result)
}

func main() {
	result1 := solution("hello", 1, 2)
	fmt.Println(result1)

	result2 := solution("I love you", 3, 6)
	fmt.Println(result2)
}
