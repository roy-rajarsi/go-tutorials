package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(workerId int, taskQueue <- chan int, processedQueue chan <- int) {
	for task := range(taskQueue) {

		// The Worker constantly listens to the Task Queue channel.
		// Since all the workers are listening at the same time, on the taskQueue,
		// only one worker can poll a task from the channel, at a time !
		// Once the task is polled, the goroutine, performs the task/job, and pushes it to
		// the Processed Queue, and again starts listening !

		fmt.Printf("Worker_%d is executing %d\n", workerId, task)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		processedQueue <- 2 * task
	}
}

func main() {
	taskCount := 10
	taskQueue := make(chan int, taskCount)
	processedQueue := make(chan int, taskCount)

	// Initiate Workers
	workerCount := 4
	for workerId := range(workerCount) {
		go worker(workerId, taskQueue, processedQueue)
		fmt.Printf("Worker_%d is initiated\n", workerId)
	}
	fmt.Println()

	// Add tasks to the TaskQueue
	for task := range(taskCount) {
		taskQueue <- task
	}
	close(taskQueue)

	// Poll from the Processed Queue
	for range(taskCount) {
		fmt.Println("Processed Task: ", <- processedQueue)
	}
}