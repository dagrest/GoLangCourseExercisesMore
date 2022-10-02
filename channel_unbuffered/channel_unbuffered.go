package main

import (
	"fmt"
)

func main() {
	fmt.Println("Channel example.")

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Put")
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
			fmt.Println("Get")
		}
	}()

	//time.Sleep(time.Second)
	fmt.Scanln()
}
