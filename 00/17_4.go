// OX퀴즈
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(quiz []string) []string {
	result := []string{}

	for _, v := range quiz {
		array := strings.Split(v, " ")
		n1, _ := strconv.Atoi(array[0])
		answer := n1

		for i := 1; i < len(array); i += 2 {
			switch array[i] {
			case "+":
				n2, _ := strconv.Atoi(array[i+1])
				answer += n2
			case "-":
				n3, _ := strconv.Atoi(array[i+1])
				answer -= n3
			case "=":
				n4, _ := strconv.Atoi(array[i+1])

				if answer == n4 {
					result = append(result, "O")
				} else {
					result = append(result, "X")
				}
			}
		}
	}

	return result
}

func main() {
	result1 := solution([]string{"3 - 4 = -3", "5 + 6 = 11"})
	fmt.Println(result1)

	result2 := solution([]string{"19 - 6 = 13", "5 + 66 = 71", "5 - 15 = 63", "3 - 1 = 2"})
	fmt.Println(result2)
}
