// 모스부호 (1)
package main

import (
	"fmt"
	"strings"
)

func solution(letter string) string {
	morse := map[string]string{
		".-":   "a",
		"-...": "b",
		"-.-.": "c",
		"-..":  "d",
		".":    "e",
		"..-.": "f",
		"--.":  "g",
		"....": "h",
		"..":   "i",
		".---": "j",
		"-.-":  "k",
		".-..": "l",
		"--":   "m",
		"-.":   "n",
		"---":  "o",
		".--.": "p",
		"--.-": "q",
		".-.":  "r",
		"...":  "s",
		"-":    "t",
		"..-":  "u",
		"...-": "v",
		".--":  "w",
		"-..-": "x",
		"-.--": "y",
		"--..": "z",
	}

	result := ""
	slice := strings.Split(letter, " ")

	for _, v := range slice {
		result += morse[v]
	}

	return result
}

func main() {
	result1 := solution(".... . .-.. .-.. ---")
	fmt.Println(result1)

	result2 := solution(".--. -.-- - .... --- -.")
	fmt.Println(result2)
}
