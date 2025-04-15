package main

import (
	"fmt"
	"math/rand"
	"time"
)

func execute(message string, taskCount int, channel chan int) {

	for taskId := 1; taskId <= taskCount; taskId ++ {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		fmt.Println("Goroutine :: ", message, taskId)
		channel <- taskId
	}
}

func main() {
	tasksExecuted := 0
	taskCount := 10
	channel := make(chan int)
	message := "Completed Execution of Task Id: "
	go execute(message, taskCount, channel)

	for tasksExecuted < taskCount {
		lastCompletedTaskId := <- channel
		tasksExecuted++
		fmt.Println("Main :: Received from channel:", lastCompletedTaskId)
	}
	fmt.Println("Main :: All Tasks Executed ....")
}
