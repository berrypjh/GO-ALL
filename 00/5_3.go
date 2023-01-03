// 나이 출력
package main

import "fmt"

func solution(age int) int {
    return 2022 - age + 1
}

func main() {
	result1 := solution(40)
	fmt.Println(result1)

	result2 := solution(23)
	fmt.Println(result2)
}