package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    c := 'a'
    switch {
    case '0' <= c && c <= '9':
        fmt.Printf("%c은(는) 숫자입니다.", c)
    case 'a' <= c && c <= 'z':
        fmt.Printf("%c은(는) 소문자입니다.", c)
    case 'A' <= c && c <= 'Z':
        fmt.Printf("%c은(는) 대문자입니다.", c)
    }
}