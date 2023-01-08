// 개미 군단
package main

import (
	"fmt"
)

func solution(hp int) int {
	result := 0

	n1 := hp % 5
	result += hp / 5
	n2 := n1 % 3
	result += n1 / 3
	result += n2 / 1

	return result
}

func main() {
	result1 := solution(23)
	fmt.Println(result1)

	result2 := solution(24)
	fmt.Println(result2)

	result3 := solution(999)
	fmt.Println(result3)
}
