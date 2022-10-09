package main

import "fmt"

var x int

func main() {
	// 클로저 : 변수의 범위를 닫고 특정 영역에 담는 것
	fmt.Println(x)
	x++

	{
		y := 32
		fmt.Println(y)
	}
	// fmt.Println(y) - 코드 블록 범위 밖이라 컴파일 에러 발생

	fmt.Println(x)
	foo()
	fmt.Println(x)

	// 아래 a 와 b 안의 x 는 범위 다르고 저장되는 메모리 위치 다름
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() {
	fmt.Println("hello")
	x++ // x 의 범위가 이 전체 패키지라 함수를 실행하면 1 씩 증가합니다.
}

func incrementor() func() int {
	var x int // 초깃값 0 이며 incrementor 함수안에서만 존재함
	return func() int {
		x++
		return x
	}
}
