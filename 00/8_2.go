// 외계행성의 나이
package main

import (
	"fmt"
	"strconv"
)

func solution(age int) string {
	result := ""

	for _, v := range strconv.Itoa(age) {
		result += string(v + 49)
	}

	return result
}

func main() {
	result1 := solution(1200)
	fmt.Println(result1)

	result2 := solution(51)
	fmt.Println(result2)

	result3 := solution(100)
	fmt.Println(result3)
}
