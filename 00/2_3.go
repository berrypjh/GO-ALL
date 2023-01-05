// 배열 두배 만들기
package main

import "fmt"

func solution(numbers []int) []int {
	// x := []int{}

	// for _, v := range numbers {
	// 	x = append(x, v * 2)
	// }

	// return x
	for i, v := range numbers {
		numbers[i] = v * 2
		fmt.Println(numbers)
	}

	return numbers
}

func main() {
	result1 := solution([]int{1, 2, 3, 4, 5})
	fmt.Println(result1)

	result2 := solution([]int{1, 2, 100, -99, 1, 2, 3})
	fmt.Println(result2)
}
