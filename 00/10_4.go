// 배열 회전시키기
package main

import "fmt"

func solution(numbers []int, direction string) []int {
	switch direction {
	case "right":
		numbers = append(numbers[len(numbers)-1:], numbers[:len(numbers)-1]...)
	case "left":
		numbers = append(numbers[1:], numbers[0])
	}
	return numbers
}

func main() {
	result1 := solution([]int{1, 2, 3}, "right")
	fmt.Println(result1)

	result2 := solution([]int{4, 455, 6, 4, -1, 45, 6}, "left")
	fmt.Println(result2)
}
