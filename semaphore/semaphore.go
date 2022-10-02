package main

import (
	"fmt"
)

func main() {

	fmt.Println("Semaphore example.")

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Put")
			c <- i
		}
		fmt.Println("Put Done")
		done <- true
	}()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Println("Put---2")
			c <- i
		}
		fmt.Println("Put---2 Done")
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
