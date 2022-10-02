package main

import "fmt"

func main() {
	n, _ := fmt.Println("Hello world", 42)
	fmt.Println(n)

	x := 42 // 변수 선언과 값을 동시 지정
	fmt.Println(x)
	x = 99 // 수정가능
	fmt.Println(x)

	y := 100 + 24
	fmt.Println(y)

	var z = 43 // 함수 본문 바깥에서 선언가능, (:= 의 경우 불가능)
	// 하지만 변수의 범위를 최대한 제한하는 것이 좋으며, 최대한 단축 선언하는 것이 좋습니다.
	fmt.Println(z)

	var a int
	fmt.Println(a) // 0

	var b string
	fmt.Println(b) // ""
}
