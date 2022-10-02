package main

import (
	"fmt"
	"time"
)

func main() {

	var name string
	fmt.Println("Channel channel_many_to_one example.")

	c := make(chan int, 4000)

	go func() {
		for i := 0; i < 2000; i++ {
			fmt.Println("Put")
			c <- i
		}
	}()

	go func() {
		for i := 2000; i < 4000; i++ {
			fmt.Println("Put---2")
			c <- i
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	close(c)

	for n := range c {
		fmt.Println(n)
	}
}
