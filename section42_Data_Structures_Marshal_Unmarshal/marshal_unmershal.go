package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int `json:"age"` // set int `json:"-"` to hide the value
}

func (p Person) printPersonFirstLast() {
	fmt.Println("Person details: ", p.First, p.Last)
}

func main() {
	fmt.Println("Marshal/Unmarshal Examples:")

	p1 := Person{"James", "Bond", 35}
	p2 := Person{"Miss", "Moneypenny", 18}

	fmt.Println(p1, p2)
	p1.printPersonFirstLast()
	p2.printPersonFirstLast()

	jsonObject1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshal action result: ", jsonObject1)
	fmt.Printf("%T \n", jsonObject1)
	fmt.Println("Marshal action result, after converting to string: ", string(jsonObject1))

	var newP1 Person
	json.Unmarshal(jsonObject1, &newP1)
	fmt.Println("Unmarshal action result: ", newP1)
	fmt.Printf("%T \n", newP1)
	//fmt.Println("Marshal action result, after converting to string: ", string(jsonObject1))

	fmt.Printf("-------\n")
	var newP2 Person
	bytesSlice := []byte(`{"First":"Miss", "Last":"Moneypenny", "Age":25}`)
	fmt.Println(bytesSlice)
	fmt.Println(string(bytesSlice))
	json.Unmarshal(bytesSlice, &newP2)
	fmt.Println("Unmarshal action result: ", newP2)
	fmt.Printf("%T \n", newP2)
}
