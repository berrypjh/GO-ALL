// 잘라서 배열로 저장하기
package main

import (
	"fmt"
)

func solution(my_str string, n int) []string {
	result := []string{}
	str := ""

	for i := 1; i <= len(my_str); i++ {
		str += string(my_str[i-1])

		if i%n == 0 {
			result = append(result, str)
			str = ""
		}
	}

	if str != "" {
		result = append(result, str)
	}

	return result
}

func main() {
	result1 := solution("abc1Addfggg4556b", 6)
	fmt.Println(result1)

	result2 := solution("abcdef123", 3)
	fmt.Println(result2)
}
