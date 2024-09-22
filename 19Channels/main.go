package main

import (
	"fmt"
	"sync"
)

func main() {

	// create a channel of type int, add a buffer of size 2
	myChannel := make(chan int, 2)

	// send values to the channel using the <- operator
	// myChannel <- 1
	// myChannel <- 2
	// myChannel <- 3

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// <- means receive ONLY from the channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// receive values from the channel using the <- operator
		fmt.Println(<-ch)
		// here channel has two listeners
		// fmt.Println(<-ch)
		val, isopen := <-ch
		fmt.Println(val, isopen)
		wg.Done()
	}(myChannel, wg)

	// <- means send ONLY to the channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// receive values from the channel using the <- operator
		ch <- 4
		ch <- 5
		wg.Done()
	}(myChannel, wg)

	wg.Wait()

	// close the channel
	defer close(myChannel)

}
