// 최빈값 구하기
package main

import "fmt"

func solution(array []int) int {
	x := make([]int, 100, 100)

	for _, v := range array {
		x[v]++
	}

	num := 0
	result := 0
	for i, v := range x {
		if num < v {
			num = v
			result = i
		}
	}

	z := 0
	for _, v := range x {
		if num == v {
			z++
		}
	}

	if z > 1 {
		return -1
	}

	return result
}

func main() {
	result1 := solution([]int{5, 5, 3, 5, 3, 4})
	fmt.Println(result1)

	result2 := solution([]int{1, 1, 2, 2})
	fmt.Println(result2)

	result3 := solution([]int{1})
	fmt.Println(result3)
}
