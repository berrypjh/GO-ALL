package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	// speak 함수는 secretAgent 타입에 연결됩니다. 해당 타입으로 이 메소드를 엑세스할 수 있습니다.
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Money",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak() // 이처럼 메소드를 엑세스할 수 있습니다.
	sa2.speak()
}
