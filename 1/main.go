package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomthing()

	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomthing() (string, string) {
	return "something", "else"
}
