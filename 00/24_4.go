// k의 개수
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(i int, j int, k int) int {
	result := 0
	str := strconv.Itoa(k)

	for x := i; x <= j; x++ {
		arr := strings.Split(strconv.Itoa(x), "")

		for _, v := range arr {
			if strings.Contains(v, str) {
				result++
			}
		}
	}

	return result
}

func main() {
	result1 := solution(1, 13, 1)
	fmt.Println(result1)

	result2 := solution(10, 50, 5)
	fmt.Println(result2)

	result3 := solution(3, 10, 2)
	fmt.Println(result3)
}
