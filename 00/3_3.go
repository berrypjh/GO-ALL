// 짝수는 싫어요
package main

import "fmt"

func solution(n int) []int {
	x := []int{}

	// for i := 1; i <= n; i+=2 {
	// 	x = append(x, i)
	// }

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			x = append(x, i)
		}
	}

	return x
}

func main() {
	result1 := solution(10)
	fmt.Println(result1)

	result2 := solution(15)
	fmt.Println(result2)
}
