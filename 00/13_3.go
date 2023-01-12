// 중복된 문자 제거
package main

import (
	"fmt"
)

func solution(my_string string) string {
	result := ""
	m := map[rune]bool{}

	for _, v := range my_string {
		_, ok := m[v]

		if !ok {
			m[v] = true
			result += string(v)
		}
	}

	return result
}

func main() {
	result1 := solution("people")
	fmt.Println(result1)

	result2 := solution("We are the world")
	fmt.Println(result2)
}
