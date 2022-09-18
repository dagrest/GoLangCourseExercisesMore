package main

import "fmt"

func visit(array []int, callbackFunc func(int)) {

	for _, v := range array {
		callbackFunc(v)
	}
}

func main() {
	fmt.Println("Callbacks")

	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})
}
