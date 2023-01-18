// 다항식 더하기
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(polynomial string) string {
	result := ""
	xnum, num := 0, 0
	array := strings.Split(polynomial, " ")

	for i := 0; i < len(array); i += 2 {
		if strings.Contains(array[i], "x") {
			xstr := strings.Trim(array[i], "x")

			if xstr == "" {
				xnum++
			} else {
				n1, _ := strconv.Atoi(xstr)
				xnum += n1
			}
		} else {
			n2, _ := strconv.Atoi(array[i])
			num += n2
		}
	}

	if xnum != 0 {
		sd := ""
		if xnum != 1 {
			sd = strconv.Itoa(xnum)
		}
		result = sd + "x"
	}
	if xnum != 0 && num != 0 {
		result += " + "
	}
	if num != 0 {
		sd := strconv.Itoa(num)
		result += sd
	}
	if xnum == 0 && num == 0 {
		result += "0"
	}

	return result
}

func main() {
	result1 := solution("3x + 7 + x")
	fmt.Println(result1)

	result2 := solution("x + x + x")
	fmt.Println(result2)

	result3 := solution("12x + x + 3x + 20")
	fmt.Println(result3)

	result4 := solution("x + 20")
	fmt.Println(result4)
}
