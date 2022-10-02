package main

import (
	"fmt"
)

func main() {
	fmt.Println("Channel range example.")

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Put")
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
