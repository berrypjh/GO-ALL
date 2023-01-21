package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// bcrypt - 비밀번호 정보를 저장하는 좋은 방식
	// .../golang.org/x/crypto/bcrypt 에 있습니다.
	// 여기서 x 란 마이그레이션 하지 않은 표준 실험 라이브러리입니다.
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password1234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}
