package main

import "fmt"

func main() {
	defer foo()
	bar()
	// foo 는 main 함수가 종료될 때 defer 되고 실행이 됩니다.
	// 만약 파일을 열고 그 위치의 작업이 끝나는 부분에 파일이 닫히길 원할 때
	// 파일연 직후 파일 닫는 함수를 defer 하여 메모리도 아끼고 가독성을 높여줍니다.
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
