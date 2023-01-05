// 특정 문자 제거하기
package main

import "fmt"

func solution(my_string string, letter string) string {
	result := ""

	for _, v := range my_string {
		if string(v) != letter {
			result += string(v)
		}
	}

	return result
}

func main() {
	result1 := solution("abcdef", "f")
	fmt.Println(result1)

	result2 := solution("BCBdbe", "B")
	fmt.Println(result2)
}
