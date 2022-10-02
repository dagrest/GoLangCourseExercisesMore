package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i <= 9; i++ {
			c <- i + 1
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}
