package main

import "fmt"

func main() {
	// 포인터 : 값이 저장된 어떤 메모리의 위치를 가리키는 것
	a := 23
	fmt.Println(a)
	fmt.Println(&a) // & : 메모리에서 주소를 보는 방법

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // *int : 주소의 타입은 int 에 대한 포인터

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // * 를 사용하여 해당 주소의 값을 볼 수 있습니다.
	fmt.Println(*&a)
	*b = 53
	fmt.Println(a)

	// 1. 포인터는 큰 데이터 덩어리가 있고 프로그램을 통해 큰 데이터 덩어리를 전달하지 않는 경우에 좋습니다.
	// 이때, 데이터가 저장된 주소를 전달하기만 하면 됩니다.
	// 데이터가 엄청 많고 데이터베이스에서 데이터를 가져오고 있다면 메모리에 저장되어 있으니
	// 그냥 주소만 전달하면 됩니다.
	// 2. 특정 위치에 있는 무언가를 변경해야하는 경우에도 포인터를 사용합니다.

	// * Go 언어의 모든 것은 값으로 전달됩니다.
	x := 0
	foo(x)
	fmt.Println(x)

	y := 0
	bar(&y)
	fmt.Println(y)
}

func foo(z int) {
	fmt.Println(z)
	z = 43
	fmt.Println(z)
}

func bar(z *int) {
	fmt.Println(z)
	*z = 43
	fmt.Println(*z)
}
