package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var input string

var wg sync.WaitGroup

const iterNum = 20

var counter int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all of the cores!!!
}

func main() {
	fmt.Println("Concurrency example 1")
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	wg.Add(2)
	go goo("Goo:")
	go goo("Foo:")
	wg.Wait()
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("Num of goroutines: ", runtime.NumGoroutine())
}

func goo(name string) {
	for i := 0; i <= iterNum; i++ {
		x := counter
		x++
		fmt.Println(name, ": ", i, ", num of goroutines: ", runtime.NumGoroutine(), ", Counter: ", counter)
		time.Sleep(time.Duration(3 * time.Millisecond))
		counter = x
	}
	wg.Done()
}
