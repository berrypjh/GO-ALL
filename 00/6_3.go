// 문자열 뒤집기
package main

import "fmt"

func solution(num_list []int) []int {
	odd, even := 0, 0
	for _, v := range num_list {
		switch v % 2 {
		case 0:
			odd++
		case 1:
			even++
		}
	}

	return []int{odd, even}
}

func main() {
	result1 := solution([]int{1, 2, 3, 4, 5})
	fmt.Println(result1)

	result2 := solution([]int{1, 3, 5, 7})
	fmt.Println(result2)
}
