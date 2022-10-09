package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64 // area() float64 메소드가 첨부된 모든 메소드를 의미
}

// func (c circle) area() float64 {
func (c *circle) area() float64 {
	// circle 은 일반 타입
	// t *T 처럼 포인터 리시버는 포인터인 값 (*T) 에서만 작동
	// t T 처럼 포인터가 아닌 리시버는 포인터나 포인터가 아닌 값 (T, *T) 으로 작동
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

type person struct {
	name string
}

func main() {
	// 메소드 세트 : 타입에 연결하는 메소드
	// 메소드 세트는 타입이 첨부되는 메소드를 결정합니다.
	// 포인터가 아닌 리시버가 있는데 리시버가 있는 타입에 함수를 첨부합니다.

	c := circle{5}
	fmt.Printf("%T\n", &c) // *main.circle
	// info(c) - (c *circle) 일때는 작동안함
	info(&c)

	p1 := person{
		name: "Tom",
	}
	fmt.Println(p1.name)
	changeMe(&p1)
	fmt.Println(p1.name)
}

func changeMe(x *person) {
	// x.name = "Kim"
	fmt.Println(x)
	(*x).name = "Kim"
}
