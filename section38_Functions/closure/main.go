package main

import "fmt"

func wrapper() func() int {
	x := 0 // more idiomatic way to do init: var x int
	//var x int

	return func() int {
		x++
		return x
	}
}

func main() {
	fmt.Println("Closure example")
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
