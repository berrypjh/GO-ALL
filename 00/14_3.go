// 암호 해독
package main

import (
	"fmt"
)

func solution(cipher string, code int) string {
	result := ""

	for i := code; i <= len(cipher); i+=code {
		result += string(cipher[i - 1])
	}

    return result
}

func main() {
	result1 := solution("dfjardstddetckdaccccdegk", 4)
	fmt.Println(result1)

	result2 := solution("pfqallllabwaoclk", 2)
	fmt.Println(result2)
}
