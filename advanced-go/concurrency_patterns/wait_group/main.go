package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerId int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()  // Once the Goroutine is done drop the WaitGroup Counter by 1
	fmt.Printf("Worker_%d started\n", workerId)
	time.Sleep(5 * time.Second)
	fmt.Printf("Worker_%d finising\n", workerId)
}

func main() {
	waitGroup := &sync.WaitGroup{}
	workerCount := 5

	for workerId := range(workerCount) {
		waitGroup.Add(1)  // Since main() must wait for this goroutine add 1 to the WaitGroup counter
		go worker(workerId, waitGroup)
	}
	waitGroup.Wait()  // main() waits so long as the WaitGroup counter is not zero
	fmt.Println("All Goroutines completed !")
}