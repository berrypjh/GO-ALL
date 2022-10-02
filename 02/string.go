package main

import "fmt"

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// 문자열은 byte(int8)로 이루졌고 생성되면 내용을 변경할 수 없습니다.
	bs := []byte(s) // 문자열 타입에서 byte 로 된 슬라이스로 변환
	fmt.Println(bs) // 출력결과 ASCII 참고
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i]) // UFT-8 코드를 줍니다.
		// 출력결과 U+0048 'H' ..... (UTF-8 코드 포인트 '해당 문자')
		// 이 코드 포인트를 rune(int32) 이라고 부름
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("index: %d, byte: %v, hex: %#x\n", i, v, v)
	}
}
