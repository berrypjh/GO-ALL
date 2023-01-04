// 문자열 뒤집기
package main

import "fmt"

func solution(my_string string) string {
	result := ""

	for i := len(my_string) - 1; i >= 0; i-- {
		result += string(my_string[i])
	}

    return result
}

func main() {
	result1 := solution("jaron")
	fmt.Println(result1)

	result2 := solution("bread")
	fmt.Println(result2)
}