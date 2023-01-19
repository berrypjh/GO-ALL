// 안전지대
package main

import (
	"fmt"
)

func solution(board [][]int) int {
	length := len(board)
	result := 0

	array := make([][]bool, length)
	for i := 0; i < length; i++ {
		array[i] = make([]bool, length)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if board[i][j] == 1 {
				array[i][j] = true

				for x := -1; x <= 1; x++ {
					if i+x < 0 || i+x >= length {
						continue
					}
					for y := -1; y <= 1; y++ {
						if j+y < 0 || j+y >= length {
							continue
						}

						array[i+x][j+y] = true
					}
				}
			}
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if !array[i][j] {
				result++
			}
		}
	}

	return result
}

func main() {
	result1 := solution([][]int{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 0}})
	fmt.Println(result1)

	result2 := solution([][]int{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 1, 1, 0}, []int{0, 0, 0, 0, 0}})
	fmt.Println(result2)

	result3 := solution([][]int{[]int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}})
	fmt.Println(result3)
}
