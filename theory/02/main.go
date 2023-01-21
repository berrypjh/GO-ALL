package main

import (
	"fmt"
	"runtime"
)

var x bool
var y int
var z float64

const a = 42 // 상수
const (
	b        = 4.2665
	c string = "main"
)

const (
	d = iota
	e
	f
)

const ( // const 키워드 사용하면 iota 0부터 다시 시작
	g = iota
	h
)

const (
	_ = iota
	// kb = 1024
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	y = 3
	z = 5.2321
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)

	// runtime 패키지
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
