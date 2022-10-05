package main

import (
	"fmt"
	"time"
)

const maxNum1000 int = 10

func main() {
	fmt.Println("Channel pipeline example")

	startTime := time.Now().UnixMilli()
	fmt.Println("Start time: ", startTime)
	in := insert()
	out := read(in)

	fmt.Println("Channels list len: ", len(in))

	time.Sleep(time.Second)

	var endTimeRes int64
	for v := range out {
		fmt.Println("Content of read channel", v)
		endTimeRes = v
	}
	fmt.Println("Calculation time: ", endTimeRes-int64(startTime))

}

type myChannel = chan int

func insert() chan myChannel {
	chList := make(chan myChannel, maxNum1000)
	go func() {
		for i := 1; i <= maxNum1000; i++ {
			fmt.Println("Starting a new routine with i: ", i)
			go func() {
				in := make(chan int)
				chList <- in
				fmt.Println("Should be added to channel: ", i)
				in <- i
				fmt.Println("Added to channel: ", i)
				close(in)
			}()
		}
	}()
	return chList
}

func read(in chan myChannel) chan int64 {
	out := make(chan int64)
	go func() {
		for ch := range in {
			go func() {
				v := <-ch
				fmt.Println("Value for factorial from channel: ", v)
				out <- int64(factorial(v))
			}()
		}
		out <- time.Now().UnixMilli()
		close(out)
		fmt.Println("End internal time: ", time.Now().UnixMilli())
	}()

	return out
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
