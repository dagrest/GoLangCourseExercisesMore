package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Rune Example")

	byteArray := []rune("動-死者-abc-人超に")
	for _, v := range byteArray {
		fmt.Print("Single rune value '", string(v), "', rune decimal value: ", v, " ")
		fmt.Printf("Type: %T, size: %d\n", v, unsafe.Sizeof(v))
	}

	fmt.Println("String: ", string(byteArray))
}
