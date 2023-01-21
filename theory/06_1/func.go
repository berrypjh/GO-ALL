package main

import "fmt"

func main() {
	// s := []int{1, 2, 3, 4}
	// x := sum(s...)
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is,", x)

	y := sum()
	// 인자를 전달하지 않으면 []int 가 출력되지만 len(), cap() 의 출력 값이 0 으로 뜹니다.
	// 이는 슬라이스는 존재하지만 포인터는 아무것도 가리키지 않고 nil 을 가리키는 상태입니다.
	fmt.Println("The total is,", y)
}

func sum(x ...int) int { // ...int : Lexical elements (이 타입은 int 의 슬라이스)
	// 인자 2개 이상일떄, 가변 인자는 뒤에 적습니다. - (s string, x ...int)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	return sum
}
