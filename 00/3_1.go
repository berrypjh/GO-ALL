// 중앙값 구하기
package main

import (
	"fmt"
	"sort"
)

func solution(array []int) int {
	/* // 정렬하기
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array) - 1; j++ {
			if array[j] > array[j + 1] {
				temp := array[j]
				array[j] = array[j + 1]
				array[j + 1] = temp
			}
		}
	}

	// 중앙 index 구하기
	index := len(array) / 2

    return array[index] */

	sort.Ints(array)
  
    return array[len(array)/2]
}

func main() {
	result1 := solution([]int{1, 2, 7, 10, 11})
	fmt.Println(result1)
	
	result2 := solution([]int{9, -1, 0})
	fmt.Println(result2)
}