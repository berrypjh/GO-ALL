package main

import "fmt"

func main() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
	}

	if x := 42; x == 42 { // if 문 안에서만 x 가 선언된다.
		fmt.Println("001")
	}
	// fmt.Println(x) - 에러 발생

	n := "D"
	switch n {
	case "A", "B", "C":
		fmt.Println("A ~ C")
	case "D":
		fmt.Println("D")
		fallthrough // 아래의 결과와 상관없이 실행, 안쓰는 것이 좋음
	default:
		fmt.Println("default")
	}
}
