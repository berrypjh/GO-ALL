// 점의 위치 구하기
package main

import "fmt"

func solution(dot []int) int {
	if dot[0] > 0 {
		if dot[1] > 0 {
			return 1
		}
		return 4
	} else {
		if dot[1] > 0 {
			return 2
		}
		return 3
	}
}

func main() {
	result1 := solution([]int{2, 4})
	fmt.Println(result1)

	result2 := solution([]int{-7, 9})
	fmt.Println(result2)
}
