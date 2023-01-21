package main

import "fmt"

func main() {
	var x [5]int
	// var y [6]int - 길이도 타입의 일부이므로 위와 서로 다른 타입이다.
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	// 합성 리터럴
	z := []int{3, 5, 6, 8, 12} // 슬라이스 선언과 동시에 값 초기화
	fmt.Println(z)

	a := []int{3, 4, 8, 9, 38}
	for i, v := range a { // 배열, 슬라이스에 for...range 를 사용하여 전체 요소를 반복하고 각 요소의 인덱스와 값을 얻어올 수 있다.
		fmt.Println(i, v)
	}

	fmt.Println(a)
	fmt.Println(a[1:])
	fmt.Println(a[2:4]) // 슬라이스 일부 값 추출
}
