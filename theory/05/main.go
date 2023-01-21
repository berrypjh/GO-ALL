package main

import "fmt"

type person struct {
	name string
	age  int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		name: "James",
		age:  23,
	}
	p2 := person{
		name: "Tom",
		age:  36,
	}
	fmt.Println(p1)
	fmt.Println(p2.name, p2.age)

	sa1 := secretAgent{
		person: person{
			name: "Kim",
			age:  26,
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			name: "Park",
			age:  32,
		},
		ltk: false,
	}
	fmt.Println(sa1)
	fmt.Println(sa2.name, sa2.age, sa2.ltk)
	fmt.Println(sa2.person.name, sa2.age, sa2.ltk) // sa2.person.name 은 sa2.name 로 사용가능

	// 익명 구조체 - 코드를 작은 영역에서 깨끗하게 쓰고 싶을 때 사용
	p3 := struct {
		name string
		age  int
	}{
		name: "A",
		age:  13,
	}
	fmt.Println(p3)
	fmt.Println(p3.name)
}
