// 외계어 사전
package main

import (
	"fmt"
	"strings"
)

func solution(spell []string, dic []string) int {
	for _, v := range dic {
		num := len(spell)

		for _, w := range spell {
			if strings.Contains(v, w) {
				num--
			}
		}

		if num == 0 {
			return 1
		}
	}

	return 2
}

func main() {
	result1 := solution([]string{"p", "o", "s"}, []string{"sod", "eocd", "qixm", "adio", "soo"})
	fmt.Println(result1)

	result2 := solution([]string{"z", "d", "x"}, []string{"def", "dww", "dzx", "loveaw"})
	fmt.Println(result2)

	result3 := solution([]string{"s", "o", "m", "d"}, []string{"moos", "dzx", "smm", "sunmmo", "som"})
	fmt.Println(result3)
}
