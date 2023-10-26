package main

import "fmt"

func main() {
	fmt.Println("Hello World.")

	var whatToSay string

	var i int

	whatToSay = "Goodbye, friends"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println(whatWasSaid)
	fmt.Println(theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something" "else"
}
