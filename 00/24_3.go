// A로 B 만들기
package main

import (
	"fmt"
	"sort"
)

func solution(before string, after string) int {
	str1 := []string{}
	for _, v := range before {
		str1 = append(str1, string(v))
	}
	sort.Strings(str1)

	str2 := []string{}
	for _, v := range after {
		str2 = append(str2, string(v))
	}
	sort.Strings(str2)

	for i, v := range str1 {
		if str2[i] != v {
			return 0
		}
	}

	return 1
}

func main() {
	result1 := solution("olleh", "hello")
	fmt.Println(result1)

	result2 := solution("allpe", "apple")
	fmt.Println(result2)
}
