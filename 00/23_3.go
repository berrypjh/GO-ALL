// 옹알이 (1)
package main

import (
	"fmt"
	"regexp"
)

func solution(babbling []string) int {
	result := 0
	re := regexp.MustCompile(`aya|ye|woo|ma`)

	for _, v := range babbling {
		temp := re.ReplaceAllString(v, "")

		if temp == "" {
			result++
		}
	}

	return result
}

func main() {
	result1 := solution([]string{"aya", "yee", "u", "maa", "wyeoo"})
	fmt.Println(result1)

	result2 := solution([]string{"ayaye", "uuuma", "ye", "yemawoo", "ayaa"})
	fmt.Println(result2)
}
