package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	// first string
	// last  string
	// age   int
	First string
	Last  string
	Age   int
}

func main() {
	// 라이브러리 사용하기 내용 블로그 정리
	p1 := person{
		// first: "James",
		// last:  "Bond",
		// age:   32,
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		// first: "Miss",
		// last:  "Moneypenny",
		// age:   27,
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bs)) // 구조체가 위의 주석대로 소문자로 작성되면 [{},{}] 로 출력됩니다.
}
