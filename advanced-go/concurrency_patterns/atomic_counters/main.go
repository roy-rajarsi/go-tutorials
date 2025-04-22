package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	// Lets say this counter is going to be utilised by a bunch of goroutines
	var counter int64 = 100

	// All goroutines must use these atomic functions to perform operations on the counter
	// All operations take in the reference of the counter !

	// Get the value of the counter
	fmt.Println("Counter :: ", atomic.LoadInt64(&counter))

	// Add an integer to the counter
	atomic.AddInt64(&counter, 10)
	fmt.Println("Counter :: ", atomic.LoadInt64(&counter))

	// Swap or overwrite the value of the counter with the new value 200. It returns the old value of the counter
	oldValue := atomic.SwapInt64(&counter, 200)
	fmt.Println("Counter :: ", atomic.LoadInt64(&counter), "Old Value :: ", oldValue)
	
	// Compare and Swap - Most imporant use-case
	// Takes in the old value and new value and returns, if CAS was successful
	couldSwap := atomic.CompareAndSwapInt64(&counter, 200, 1000)  // couldSwap -> True, since oldValue = 200
	fmt.Println("Counter :: ", counter, "Could Swap :: ", couldSwap)

	couldSwap = atomic.CompareAndSwapInt64(&counter, 10, 2000) // couldSwap -> False, since oldValue = 1000 and not 10
	fmt.Println("Counter :: ", counter, "Could Swap :: ", couldSwap)
}
