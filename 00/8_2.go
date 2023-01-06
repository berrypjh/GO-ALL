// 외계행성의 나이
package main

import (
	"fmt"
)

func solution(age int) string {
	result := ""

	if age/1000 != 0 {
		result += string('a' + age/1000)
		result += string('a' + (age%1000-age%100)/100)
		result += string('a' + (age%100-age%10)/10)
	} else if (age%1000-age%100)/100 != 0 {
		result += string('a' + age/100)
		result += string('a' + (age%100-age%10)/10)
	} else if (age%100-age%10)/10 != 0 {
		result += string('a' + age/10)
	}
	result += string('a' + age%10)

	return result
}

func main() {
	result1 := solution(1200)
	fmt.Println(result1)

	result2 := solution(51)
	fmt.Println(result2)

	result3 := solution(100)
	fmt.Println(result3)
}
