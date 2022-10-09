package main

import "fmt"

func main() {
	// func 는 다른 타입과 똑같음
	// 변수에 string 할당하는 것처럼 변수에 함수를 할당
	f := func() {
		fmt.Println("Hello")
	}
	f()

	g := func(x int) {
		fmt.Println(x)
	}
	g(123)
}
