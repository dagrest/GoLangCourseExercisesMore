package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	fmt.Println("Map Examples")

	p0 := Person{"James", "Bond", 37}
	p1 := Person{"Miss", "Moneypenny", 25}

	fmt.Println(p0, p1)

	m := make(map[string]Person, 5)

	m["007"] = p0
	m["mp"] = p1

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Map length: ", len(m))

	value, exists := m["007"]
	fmt.Println("Value 007: ", value, exists)

	value, exists = m["001"]
	fmt.Println("Value 001: ", value, exists)

	for s, v := range m {
		fmt.Println("Value: ", s, v)
	}
}
