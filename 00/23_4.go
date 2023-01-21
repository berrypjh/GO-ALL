// 로그인 성공?
package main

import (
	"fmt"
)

func solution(id_pw []string, db [][]string) string {
	return ""
}

func main() {
	result1 := solution([]string{"meosseugi", "1234"}, [][]string{{"rardss", "123"}, {"yyoom", "1234"}, {"meosseugi", "1234"}})
	fmt.Println(result1)

	result2 := solution([]string{"programmer01", "15789"}, [][]string{{"programmer02", "111111"}, {"programmer00", "134"}, {"programmer01", "1145"}})
	fmt.Println(result2)

	result3 := solution([]string{"rabbit04", "98761"}, [][]string{{"jaja11", "98761"}, {"krong0313", "29440"}, {"rabbit00", "111333"}})
	fmt.Println(result3)
}
