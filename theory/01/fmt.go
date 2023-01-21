package main

import "fmt"

var y int = 90

type myType int // 타입은 myType 이고 실제 타입은 int
var a myType
var b int

func main() {
	a = 2
	b = 3
	fmt.Printf("%T\n", a) // main.myType 으로 출력됨 - 사용자만의 타입을 만들 수 있음
	// b = a - 둘은 타입이 달라 에러
	b = int(a) // conversions https://go.dev/doc/effective_go#conversions
	fmt.Println(b)

	// https://pkg.go.dev/fmt
	fmt.Println(y)
	fmt.Printf("%T\n", y)  // 타입
	fmt.Printf("%b\n", y)  // 2진수
	fmt.Printf("%x\n", y)  // 16진수
	fmt.Printf("%#x\n", y) // 0x + 16진수
	fmt.Printf("%#x\n%x\n%b\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y) // 문자열 계산하고 s 변수에 저장
	fmt.Println(s)
}
