// 문자열 계산하기
package main

import (
	"fmt"
)

func solution(s1 []string, s2 []string) int {
	m := map[string]bool{}
	result := 0

	for _, v := range s1 {
		m[v] = true
	}

	for _, v := range s2 {
		_, ok := m[v]
		if ok {
			result++
		}
	}

	return result
}

func main() {
	result1 := solution([]string{"a", "b", "c"}, []string{"com", "b", "d", "p", "c"})
	fmt.Println(result1)

	result2 := solution([]string{"n", "omg"}, []string{"m", "dot"})
	fmt.Println(result2)
}
