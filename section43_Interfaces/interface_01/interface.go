package main

import (
	"fmt"
	"os"
)

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) printPersonFirstLast() {
	fmt.Println("Person details: ", p.first, p.last)
}

type Actor struct {
	first   string
	last    string
	isActor bool
}

func (p Actor) printPersonFirstLast() {
	fmt.Println("Actor details: ", p.first, p.last)
}

type Human interface {
	printPersonFirstLast()
}

func isHuman(h Human) {
	h.printPersonFirstLast()
}

func main() {
	fmt.Println("Interface Examples:")

	p1 := Person{"James", "Bond", 35}
	p2 := Person{"Miss", "Moneypenny", 18}

	fmt.Println(p1, p2)
	p1.printPersonFirstLast()
	p2.printPersonFirstLast()

	s1 := Actor{"Bill", "Murray", true}

	fmt.Println(" ---> Interface func:")
	isHuman(p1)
	isHuman(s1)

	//var w io.Writer
	w := os.Stdout
	w.Write([]byte("Stdout -> David"))
}
