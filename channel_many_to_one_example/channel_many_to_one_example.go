package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Channel channel_many_to_one_example example.")

	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Put")
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Put---2")
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
