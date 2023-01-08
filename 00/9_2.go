// 가위 바위 보
package main

import (
	"fmt"
)

func solution(rsp string) string {
	result := ""

	for _, v := range rsp {
		switch string(v) {
		case "0":
			result += "5"
		case "2":
			result += "0"
		case "5":
			result += "2"
		}
	}

	return result
}

func main() {
	result1 := solution("2")
	fmt.Println(result1)

	result2 := solution("205")
	fmt.Println(result2)
}
