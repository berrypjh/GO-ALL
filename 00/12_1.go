// 모음 제거
package main

import "fmt"

func solution(my_string string) string {
	result := ""

	for _, v := range my_string {
		switch string(v) {
		case "a", "e", "i", "o", "u":
		default:
			result += string(v)
		}
	}

	return result
}

func main() {
	result1 := solution("bus")
	fmt.Println(result1)

	result2 := solution("nice to meet you")
	fmt.Println(result2)
}
