// 분수의 덧셈
package main

import "fmt"

func solution(denum1 int, num1 int, denum2 int, num2 int) []int {
	numerator := denum1 * num2 + denum2 * num1
	denominator := num1 * num2
    divisor := 1
    
    // 약수 구하기
    for i := 1; i <= numerator; i++ {
        if numerator % i == 0 && denominator % i == 0 {
            divisor = i
        }
    }
    
    return []int{numerator / divisor, denominator / divisor}
}

func main() {
	result1 := solution(1, 2, 3, 4)
	fmt.Println(result1)
	
	result2 := solution(9, 2, 1, 3)
	fmt.Println(result2)
}