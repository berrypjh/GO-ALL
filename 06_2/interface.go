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
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	// speak 함수는 secretAgent 타입에 연결됩니다. 해당 타입으로 이 메소드를 엑세스할 수 있습니다.
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// 키워드, 식별자, 타입 순서이며
// 인터페이스는 타입입니다.
type human interface {
	speak()
	// 즉 speak 라는 메소드를 가진 다른 타입도 type human 입니다.
	// 값은 하나 이상의 타입일 수 있습니다.

	// 이 안의 메소드(speak())에 있는 모든 타입은 human 인터페이스를 구현한다.
	// 간단하게 speak 메소드가 type human 안에 있으니 해당 타입도 된다.
}

// 아래처럼 빈 인터페이스가 있다면,
// 적어도 다른 모든 타입에는 메소드가 없을 것입니다.
// 따라서 모든 것이 비어 있는 타입 인터페이스를 구현합니다.
// type human interface {
//
// }

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrr", h.(person).first)
		// (person) : assertion - 어떤 것이 특정 타입이라고 단언하는 것
	case secretAgent:
		fmt.Println("I was passed into barrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	// 이 sa1 의 값은 type secretAgent 이며 여기에 첨부된 speak 메소드가 있기 때문에 type human 이기도 합니다.
	// 정리해서 sa1 은 type secretAgent 이면서 type human 입니다.
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

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak() // 이처럼 메소드를 엑세스할 수 있습니다.
	sa2.speak()

	fmt.Println(p1)

	// 다형성
	bar(sa1) // type secretAgent, human
	bar(sa2)
	bar(p1) // type person, human
	// 인터페이스는 값이 하나 이상의 타입이 될 수 있도록 합니다.

	// 1. conversion 복습
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x) // int(x) : conversion (전환)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// 1. assertion - 어떤 것이 특정 타입이라고 단언하는 것 (상단 참고)
}
