// 캐릭터의 좌표
package main

import (
	"fmt"
)

func solution(keyinput []string, board []int) []int {
	result := []int{0, 0}

	for _, v := range keyinput {
		switch v {
		case "left":
			if result[0] > -(board[0] / 2) {
				result[0]--
			}
		case "right":
			if result[0] < board[0]/2 {
				result[0]++
			}
		case "up":
			if result[1] < board[1]/2 {
				result[1]++
			}
		case "down":
			if result[1] > -(board[1] / 2) {
				result[1]--
			}
		}
	}

	return result
}

func main() {
	result1 := solution([]string{"left", "right", "up", "right", "right"}, []int{11, 11})
	fmt.Println(result1)

	result2 := solution([]string{"down", "down", "down", "down", "down"}, []int{7, 9})
	fmt.Println(result2)
}
