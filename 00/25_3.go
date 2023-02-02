// 연속된 수의 합
package main

import (
	"fmt"
)

func solution(num int, total int) []int {
	result := []int{}
	n := 0

	for i := -100; i <= 100; i++ {
		n1 := i + (i + (num - 1))
		n2 := n1 * num /2
		if n2 == total {
			n = i
			break
		}
	}
	fmt.Println(n)

    return result
}

func main() {
	result1 := solution(3, 12)
	fmt.Println(result1)

	result2 := solution(5, 15)
	fmt.Println(result2)

	result3 := solution(4, 14)
	fmt.Println(result3)

	result4 := solution(5, 5)
	fmt.Println(result4)
}
