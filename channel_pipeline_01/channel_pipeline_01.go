package main

import (
	"fmt"
	"time"
)

const maxNum int = 10

func main() {
	fmt.Println("Channel pipeline example")

	startTime := time.Now().UnixMilli()
	fmt.Println("Start time: ", startTime)
	in := insert()
	out := read(in)

	fmt.Println("Calculation time: ")

	var endTimeRes int64
	for v := range out {
		fmt.Println("Content of read channel", v)
		endTimeRes = v
	}
	fmt.Println("Calculation time: ", endTimeRes-int64(startTime))

}

func insert() chan int {
	in := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			for j := 3; j < maxNum+3; j++ {
				in <- j
			}
		}
		close(in)
	}()
	return in
}

func read(in chan int) chan int64 {
	out := make(chan int64)
	go func() {
		for v := range in {
			out <- int64(factorial(v))
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
