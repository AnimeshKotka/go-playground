package main

import (
	"fmt"
	"sync"
)

// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for, and then each of the goroutines
// calls Done when it is finished. At the same time, Wait can
// be used to block until all goroutines have finished.
var waitGrp sync.WaitGroup

// Mutex is a mutual exclusion lock.
// The zero value for a Mutex is an unlocked mutex.
//
// A Mutex must not be copied after first use.
var mut sync.Mutex

func main() {
	//This is usual way to run goroutine
	waitGrp.Add(2)
	go greetings("dfrd")
	// When we run the goroutine without wait group, main function
	// terminates even before the goroutine completes. So, we need
	// to add the wait group to wait until the goroutine completes.
	// If we don't use wait group, the program will terminate before
	// the goroutine completes.
	// defer greetings("aaaasj")
	// time.Sleep(3 * time.Second)

	//Using sync and wait group
	go greetings("mjjdifj")
	waitGrp.Wait()
}

func greetings(name string) {
	defer waitGrp.Done()
	mut.Lock()
	fmt.Println("Hello ", name)
	mut.Unlock()
}
