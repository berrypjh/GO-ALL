package main

import "fmt"

func main() {
	// 재귀는 함수가 자기 자신을 호출하는 것
	// 재귀로 하는 모든 작업은 루프로도 할 수 있습니다.
	// 재귀는 코드가 복잡하고 프로그램의 메모리 사용에 부정적인 영향을 미칩니다.
	n := factorial(4)
	fmt.Println(n)

	n2 := loopFact(4)
	fmt.Println(n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
