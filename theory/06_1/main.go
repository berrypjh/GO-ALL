package main

import "fmt"

func main() {
	// 함수 - 모듈화의 모든 것
	// 함수의 기본구조를 알아보자
	foo()
	bar("Tom")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

	total := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is,", total)
}

// func (r receiver) identifier(parameter(s)) (return(s)) { ... }
func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false
	return a, b
}

func sum(x ...int) int { // ...int : Lexical elements (이 타입은 int 의 슬라이스)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	return sum
}
