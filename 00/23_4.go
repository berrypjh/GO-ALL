// 로그인 성공?
package main

import (
	"fmt"
)

func solution(id_pw []string, db [][]string) string {
	return ""
}

func main() {
	result1 := solution([]string{"meosseugi", "1234"}, [][]string{[]string{"rardss", "123"}, []string{"yyoom", "1234"}, []string{"meosseugi", "1234"}})
	fmt.Println(result1)

	result2 := solution([]string{"programmer01", "15789"}, [][]string{[]string{"programmer02", "111111"}, []string{"programmer00", "134"}, []string{"programmer01", "1145"}})
	fmt.Println(result2)

	result3 := solution([]string{"rabbit04", "98761"}, [][]string{[]string{"jaja11", "98761"}, []string{"krong0313", "29440"}, []string{"rabbit00", "111333"}})
	fmt.Println(result3)
}
