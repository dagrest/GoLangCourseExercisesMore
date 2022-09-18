package main

import "fmt"

func main() {
	fmt.Println("Variadic Parameters example:")
	data := []float64{1, 2, 3, 4, 5, 6}
	n := average(data...)
	fmt.Println("Array: ", n)
	n = average(1, 2, 3, 4, 5, 6)
	fmt.Println("Variadic list: ", n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
