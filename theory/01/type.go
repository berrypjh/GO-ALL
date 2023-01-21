package main

import "fmt"

var y = 42
var z = "Shaken, not stirred"
var a string = `James said, "Shaken, not stirred"`
var b string = `James said,

"Shaken, not stirred"`

func main() {
	// Go 는 타입이 전부
	fmt.Println(y)
	fmt.Printf("%T\n", y) // 타입 출력
	fmt.Println(z)
	fmt.Printf("%T\n", z) // 타입 출력
	fmt.Println(a)
	fmt.Printf("%T\n", a) // 타입 출력
	fmt.Println(b)
	fmt.Printf("%T\n", b) // 타입 출력

	// z = 43 - string 타입이기 때문에 다른 타입인 int 형 선언 불가
	// go 언어는 STATIC programming language (not a DYNAMIC programming language)
}
