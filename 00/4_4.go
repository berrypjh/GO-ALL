// 배열의 평균값
package main

import "fmt"

func solution(numbers []int) float64 {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

    return float64(sum) / float64(len(numbers))
}

func main() {
	result1 := solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(result1)

	result2 := solution([]int{89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99})
	fmt.Println(result2)
}