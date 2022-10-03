package main

import "fmt"

func main() {
	m := map[string]int{ // map[키타입]값타입{ 초기값 }
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)
	fmt.Println(m["two"])
	fmt.Println(m["four"]) // map 에 값이 저장되어 있지 않아 0을 반환함 - 진짜 값이 0 인지 없어서 0 인지 확인 불가

	v, ok := m["four"] // 값 체크할때 대다수 ok 변수를 사용함 (관용)
	fmt.Println(v)     // 0 출력
	fmt.Println(ok)    // false 출력

	if v, ok := m["four"]; ok { // 관용적으로 많이 사용하는 코드
		fmt.Println("값 존재시 출력: ", v)
	}

	delete(m, "three") // 제거 - 참고로 없는 값도 컴파일 오류없이 제거 가능
	fmt.Println(m)
}
