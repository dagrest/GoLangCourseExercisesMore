package main

import (
	"fmt"
	"runtime"
)

var input string

//var wg sync.WaitGroup
const iterNum = 2000

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all of the cores!!!
}

func main() {
	fmt.Println("Concurrency example 1")
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	//wg.Add(2)
	go goo()
	go foo()
	//wg.Wait()
	fmt.Scan(&input)
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("Num of goroutines: ", runtime.NumGoroutine())
}

func goo() {
	for i := 0; i <= iterNum; i++ {
		fmt.Println("Goo--->: ", i, ", num of goroutines: ", runtime.NumGoroutine())
		//time.Sleep(time.Duration(3 * time.Millisecond))
	}
	//wg.Done()
}

func foo() {
	for i := 0; i <= iterNum; i++ {
		fmt.Println("Foo: ", i, ", num of goroutines: ", runtime.NumGoroutine())
		//time.Sleep(time.Duration(10 * time.Millisecond))
	}
	//wg.Done()
}
