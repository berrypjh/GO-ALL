// 이진수 더하기
package main

import (
	"fmt"
	"strconv"
)

func solution(bin1 string, bin2 string) string {
	n1, _ := strconv.ParseInt(bin1, 2, 64)
	n2, _ := strconv.ParseInt(bin2, 2, 64)

	return strconv.FormatInt(n1+n2, 2)
}

func main() {
	result1 := solution("10", "11")
	fmt.Println(result1)

	result2 := solution("1001", "1111")
	fmt.Println(result2)
}
