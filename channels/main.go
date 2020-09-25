// An unbuffered channel
package main

import (
	"fmt"
	"sync"
)

// wg is used to wait for the program to finish
var wg sync.WaitGroup

func main() {
	//declare an unbuffered channel
	count := make(chan int)
	// Add a count of 2 (one for each goroutine) to wg counter
	wg.Add(2)

	fmt.Println("Start Goroutines with the labels A and B")
	// launches 2 goroutines passing the channel
	go PrintCounts("A", count)
	go PrintCounts("B", count)

	// start channel by sending a value of 1
	fmt.Println("Channel Begin")
	count <- 1

	fmt.Println("Wait for goroutines to finsh")
	// Block program termination until goroutines finish
	wg.Wait()
	//Terminating program
	fmt.Println("\nterminating program")
}

func PrintCounts(label string, count chan int) {
	// schedule calls to wg Done which will be called when goroutine is executed and decrement the wg counter by 1
	defer wg.Done()
	for {
		// Receive msgs from channel, also blocks the receiver until msg is available into the channel
		val, ok := <-count
		if !ok {
			fmt.Println("channel was closed")
		}
		fmt.Printf("Count: %d received from %s \n\n", val, label)
		if val == 10 {
			fmt.Printf("channel closed from %s \n", label)
			close(count)
			return
		}
		val++
		// send count back to the other goroutine
		count <- val
	}
}
