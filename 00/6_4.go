// 문자 반복 출력하기
package main

import "fmt"

func solution(my_string string, n int) string {
	result := ""

	for _, v := range my_string {
		for i := 0; i < n; i++ {
			result += string(v)
		}
	}

	return result
}

func main() {
	result1 := solution("hello", 3)
	fmt.Println(result1)
}
