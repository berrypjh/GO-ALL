package main

import (
	"fmt"
	"sort"
)

func main() {
	// 슬라이스를 정렬해보겠습니다.
	xi := []int{4, 7, 2, 24, 55, 98, 34, 23, 1, 5}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr", "A"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
