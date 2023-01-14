// 한 번만 등장한 문자
package main

import (
	"fmt"
	"sort"
	"strings"
)

func solution(s string) string {
	m := map[rune]bool{}
	result := []string{}

	for _, v := range s {
		_, ok := m[v]

		if !ok {
			m[v] = true
		} else {
			m[v] = false
		}
	}

	for i, b := range m {
		if b {
			result = append(result, string(i))
		}
	}

	sort.Strings(result)

	return strings.Join(result, "")
}

func main() {
	result1 := solution("abcabcadc")
	fmt.Println(result1)

	result2 := solution("abdc")
	fmt.Println(result2)

	result3 := solution("hello")
	fmt.Println(result3)
}
