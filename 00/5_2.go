// 아이스 아메리카노
package main

import "fmt"

func solution(money int) []int {
	return []int{money / 5500, money % 5500}
}

func main() {
	result1 := solution(5500)
	fmt.Println(result1)

	result2 := solution(15000)
	fmt.Println(result2)
}
