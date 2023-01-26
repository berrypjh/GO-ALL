// 치킨 쿠폰
package main

import (
	"fmt"
)

func solution(chicken int) int {
	result := 0
	coupon := chicken

	for coupon >= 10 {
		x := coupon / 10
		result += x

		coupon = coupon%10 + x
	}

	return result
}

func main() {
	result1 := solution(100)
	fmt.Println(result1)

	result2 := solution(1081)
	fmt.Println(result2)

	result3 := solution(109)
	fmt.Println(result3)
}
