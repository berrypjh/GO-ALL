package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person

// sort 는 이러한 메소드가 주어지고 이 메소드를 사용해서 정렬하고,
// 3 개의 메소드는 소스에서 해당 인터페이스를 구현합니다.
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	// 나이 혹은 이름 정렬하기
	fmt.Println(people)
	sort.Sort(ByAge(people)) // sort.Sort 는 Interface 를 요청합니다. - func Sort(data Interface)
	fmt.Println(people)
}
