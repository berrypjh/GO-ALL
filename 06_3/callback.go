package main

import "fmt"

func main() {
	// 콜백은 인수로 func 를 전달하는 것을 의미
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println(s)

	s2 := even(sum, ii...)
	fmt.Println(s2)
}

func sum(xi ...int) int { // type int 의 가변 매개변수를 받아 들어오면 int 슬라이스가 됨 (복습)
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// 짝수 또는 홀수만 실행하기
func even(f func(xi ...int) int, vi ...int) int {
	// 함수 func(xi ...int) int 는 xi 에 할당된 int 의 가변 매개변수를 받아 int 로 반환하는(콜백) 함수입니다.
	// 이 함수는 f 변수에 할당됩니다. func(xi ...int) 표현식은 변수에 함수를 할당하는 것처럼 변수를 할당하고 있습니다.
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}
