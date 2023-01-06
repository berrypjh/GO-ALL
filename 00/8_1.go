// 배열 자르기
package main

import "fmt"

func solution(numbers []int, num1 int, num2 int) []int {
	return numbers[num1 : num2+1]
}

func main() {
	result1 := solution([]int{1, 2, 3, 4, 5}, 1, 3)
	fmt.Println(result1)

	result2 := solution([]int{1, 3, 5}, 1, 2)
	fmt.Println(result2)
}
