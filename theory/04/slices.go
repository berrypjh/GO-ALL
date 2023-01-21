package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 32, 52}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 100) // 추가하기
	fmt.Println(x)

	y := []int{243, 356, 467, 799}
	x = append(x, y...)
	fmt.Println(x)
	x = append(x[:2], x[4:]...) // 값 제거하기
	fmt.Println(x)

	// 만들어야할 총 슬라이스 크기를 미리 안다면 make 를 이용할 수 있습니다.
	z := make([]int, 10, 11) //make(타입, 길이, 용량) - 슬라이스 길이는 10 인데 하부의 배열 100개의 자리가 있음
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z = append(z, 11)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z = append(z, 12) // 용량이 초과되어 이전 배열을 버리고 새로 늘어난 하부 배열을 반환
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	as := []string{"A", "B", "C"}
	bs := []string{"a", "b", "c"}
	result := [][]string{as, bs}
	fmt.Println(result)
}
