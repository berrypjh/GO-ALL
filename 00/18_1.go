// 문자열안에 문자열
package main

import (
	"fmt"
	"strings"
)

func solution(str1 string, str2 string) int {
	if strings.Contains(str1, str2) {
		return 1
	}

    return 2
}

func main() {
	result1 := solution("ab6CDE443fgh22iJKlmn1o", "6CD")
	fmt.Println(result1)

	result2 := solution("ppprrrogrammers", "pppp")
	fmt.Println(result2)

	result3 := solution("AbcAbcA", "AAA")
	fmt.Println(result3)
}
