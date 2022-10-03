package main

import "fmt"

func main() {
	x := make([]string, 3, 3)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, ` One`, ` Two`, ` Three`) // 기존에 다시 할당되어 원하는 값이 안나옴

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	fmt.Println("")

	// make 는 아래처럼
	y := make([]string, 3, 3)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	states := []string{` One`, ` Two`, ` Three`}
	for i, v := range states {
		y[i] = v
	}

	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	for i := 0; i < len(y); i++ {
		fmt.Println(i, y[i])
	}
}
