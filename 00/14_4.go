// 대문자와 소문자
package main

import (
	"fmt"
)

func solution(my_string string) string {
	result := ""

	for _, v := range my_string {
		if 'a' <= v && v <= 'z' {
			result += string(v - 32)
		} else if 'A' <= v && v <= 'Z' {
			result += string(v + 32)
		}
	}

	return result
}

func main() {
	result1 := solution("cccCCC")
	fmt.Println(result1)

	result2 := solution("abCdEfghIJ")
	fmt.Println(result2)
}
