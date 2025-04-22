package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longRunningTask(done chan <- bool) {
	timeNeeded := time.Duration(rand.Intn(10e3)) * time.Millisecond 
	fmt.Println("Time needed for long running task: ", timeNeeded)
	time.Sleep(timeNeeded)
	done <- true
}


func main() {
	timeout := time.Duration(5 * time.Second)
	timer := time.NewTimer(timeout)
	fmt.Println("Timer started at ", time.Now())

	done := make(chan bool)
	go longRunningTask(done)

	select {
	case <- done:
		fmt.Println("Long Running Task Completed...")
	case <- timer.C:
		fmt.Println("Long Running Task Timed out...")
	}
}