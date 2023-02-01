// 문자열 밀기
package main

import (
	"fmt"
)

func solution(A string, B string) int {
	result := 0

	for i := range A {
		str := ""
		str += string(A[len(A)-i:])
		str += string(A[:len(A)-i])

		if str == B {
			break
		}
		result++
	}

	if len(A) == result {
		return -1
	}
	return result
}

func main() {
	result1 := solution("hello", "ohell")
	fmt.Println(result1)

	result2 := solution("apple", "elppa")
	fmt.Println(result2)

	result3 := solution("atat", "tata")
	fmt.Println(result3)

	result4 := solution("abc", "abc")
	fmt.Println(result4)
}
