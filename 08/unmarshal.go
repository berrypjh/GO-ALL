package main

import (
	"encoding/json"
	"fmt"
)

// JSON to GO 에서 얻은 자료구조
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// Unmarshal 은 프로그램에 들어오는 데이터를 취해서
	// 그것을 Go 자료 구조로 변환하고 그것을 모두 Go 자료 구조에 입력해야 합니다.

	// main.go 로 얻어낸 JSON 을 이용합니다.
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	// 문서에 나온대로 바이트 슬라이스로 변환해줍니다.
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	fmt.Println(bs)
	fmt.Println(string(bs))

	// people := []person{}
	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nall of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
