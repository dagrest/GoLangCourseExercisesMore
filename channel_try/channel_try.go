package main

import (
	"fmt"
	"sync"
)

func main() {

	maxNumQuantiny := 10
	var wg sync.WaitGroup
	wg.Add(maxNumQuantiny)

	var wg1 sync.WaitGroup
	wg1.Add(maxNumQuantiny)

	outIntChan := fill(maxNumQuantiny, &wg)
	resChan := read(maxNumQuantiny, outIntChan, &wg1)

	//time.Sleep(5 * time.Second)
	//fmt.Println("After sleep")

	go func(outIntChan chan int) {
		wg.Wait()
		wg1.Wait()
		fmt.Println("#####################  outIntChan CLOSED #######################")
		close(outIntChan)
		close(resChan)
	}(outIntChan)

	fmt.Println("-------------------------  Out channel length: ", len(outIntChan))
	for v := range resChan {
		fmt.Println("resChan: ", v)
	}
}

func read(maxNumQuantiny int, in chan int, wg1 *sync.WaitGroup) chan int {
	out := make(chan int)
	for i := 0; i < maxNumQuantiny; i++ {
		go func() {
			fmt.Println("read channel")
			for v := range in {
				fmt.Println("Try to insert value to read channel: ", v)
				out <- v
				fmt.Println("Inserted value to  read channel: ", v)
				wg1.Done()
			}
		}()
	}

	return out
}

func fill(maxNumQuantiny int, wg *sync.WaitGroup) chan int {

	out := make(chan int)
	for i := 0; i < maxNumQuantiny; i++ {
		go func(n int) {
			fmt.Println("Try to insert value to channel: ", n)
			out <- n
			fmt.Println("Inserted value to channel: ", n)
			wg.Done()
		}(i)
	}
	return out
}
