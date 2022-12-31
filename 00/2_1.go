// 두 수의 나눗셈
package main

import "fmt"

func solution(num1 float64, num2 float64) int {
    return int(num1 / num2 * 1000)
}

func main() {
	result := solution(3, 2)
	fmt.Println(result)
}