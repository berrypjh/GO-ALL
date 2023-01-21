package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	s2 := func() int {
		return 451
	}
	fmt.Printf("%T\n", s2) // type : func() int

	s3 := bar()
	fmt.Printf("%T\n", s3) // type : func() int
	x := s3()
	fmt.Println(x)
	fmt.Println(s3())
	fmt.Println(bar()())
}

func foo() string {
	s := "Hello world"
	return s
}

func bar() func() int { // func() int 의 type 을 반환
	return func() int {
		return 451
	}
}
