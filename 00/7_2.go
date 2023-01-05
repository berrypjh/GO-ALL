// 각도기
package main

import "fmt"

func solution(angle int) int {
	if angle < 90 {
		return 1
	} else if angle == 90 {
		return 2
	} else if angle < 180 {
		return 3
	} else if angle == 180 {
		return 4
	}

	return 0

	// if 0 < angle && angle < 90 {
	// 	result = 1
	// } else if angle == 90 {
	// 	result = 2
	// } else if 90 < angle && angle < 180 {
	// 	result = 3
	// } else if angle == 180 {
	// 	result = 4
	// }

	// return result
}

func main() {
	result1 := solution(70)
	fmt.Println(result1)

	result2 := solution(91)
	fmt.Println(result2)

	result3 := solution(180)
	fmt.Println(result3)
}
