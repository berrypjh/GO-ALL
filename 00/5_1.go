// 옷가게 할인 받기
package main

import "fmt"

func solution(price int) int {
	result := 0

	if price >= 500000 {
		result = int(float64(price) * 0.8)
	} else if  price >= 300000 {
		result = price * 9 / 10
	} else if  price >= 100000 {
		result = price * 95 / 100
	} else {
		result = price
	}

    return result
}

func main() {
	result1 := solution(150000)
	fmt.Println(result1)

	result2 := solution(580000)
	fmt.Println(result2)

	result3 := solution(10)
	fmt.Println(result3)
}