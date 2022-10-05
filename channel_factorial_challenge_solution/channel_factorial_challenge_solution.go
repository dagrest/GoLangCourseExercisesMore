package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Channel factorial Challenge example")
	startTime := time.Now().UnixMilli()
	resChan := make(chan int)
	outChan := make(chan int)
	factorial(resChan, 4, 2, 6, 5, 1, 3)
	getRes(resChan, outChan)

	endTime := time.Now().UnixMilli()
	fmt.Println("*************************** result time: ", endTime-startTime)
	time.Sleep(5 * time.Second)
	fmt.Println("After sleep")
	for n := range outChan {
		fmt.Println(n)
	}
}

func getRes(inChan chan int, outChan chan int) chan int {
	for v := range inChan {
		go func() {
			outChan <- v
		}()
	}
	close(outChan)

	return outChan
}

func factorial(resChan chan int, n ...int) {
	fmt.Println("Len: ", len(n))
	//out := make(chan int)
	for _, v := range n {
		//fmt.Println(v)
		go func() {
			total := 1
			for i := v; i > 0; i-- {
				total *= i
			}
			resChan <- total
			fmt.Println("Fact calculated for: ", v, " is: ", total)
			close(resChan)
		}()
	}
}
