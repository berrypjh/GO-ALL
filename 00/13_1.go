// 컨트롤 제트
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(s string) int {
	result := 0
	array := strings.Split(s, " ")

	for i, v := range array {
		if v == "Z" {
			n1, _ := strconv.Atoi(array[i-1])
			result -= n1
			continue
		}
		n2, _ := strconv.Atoi(v)
		result += n2
	}

	return result
}

func main() {
	result1 := solution("1 2 Z 3")
	fmt.Println(result1)

	result2 := solution("10 20 30 40")
	fmt.Println(result2)

	result3 := solution("10 Z 20 Z 1")
	fmt.Println(result3)

	result4 := solution("10 Z 20 Z")
	fmt.Println(result4)

	result5 := solution("-1 -2 -3 Z")
	fmt.Println(result5)
}
