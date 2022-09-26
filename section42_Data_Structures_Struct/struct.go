package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) printPersonFirstLast() {
	fmt.Println("Person details: ", p.first, p.last)
}

func main() {
	fmt.Println("Struct Examples:")

	p1 := Person{"James", "Bond", 35}
	p2 := Person{"Miss", "Moneypenny", 18}

	fmt.Println(p1, p2)
	p1.printPersonFirstLast()
	p2.printPersonFirstLast()
}
