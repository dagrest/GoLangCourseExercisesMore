package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Make Greeting func"
	}
}

func main() {
	greeting := func() {
		fmt.Println("Func expression")
	}

	greeting()

	makeGreeting := makeGreeter()
	fmt.Println(makeGreeting())
	fmt.Printf("'makeGreeting' has the following type: %T\n", makeGreeting)
}
