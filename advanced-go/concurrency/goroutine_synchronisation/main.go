package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct{
	Id int
	Title string
}


func execute(tasks []Task, lastCompletedTask chan Task) {
	for _, task := range(tasks) {
		fmt.Printf("Execute Goroutine :: Starting Task : %v\n", task)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		fmt.Printf("Execute Goroutine :: Task %v completed\n", task)
		lastCompletedTask <- task
	}
	close(lastCompletedTask)
	fmt.Println("Execution Completed....")
}

func logToCompletedQueue (lastCompletedTask chan Task) {
	for task := range(lastCompletedTask) {
		fmt.Printf("Log Goroutine :: Starting to log Task : %v\n", task)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		fmt.Printf("Log Goroutine :: Logged Task: %v\n", task)
	}
	fmt.Println("Logging Completed....")
}


func main() {
	tasks := []Task {
		Task {Id: 1234, Title: "Coding"},
		Task {Id: 2345, Title: "Testing"},
		Task {Id: 3456, Title: "Deploying"},
	}

	lastCompletedTask := make(chan Task)
	go execute(tasks, lastCompletedTask)
	go logToCompletedQueue(lastCompletedTask)
	time.Sleep(10 * time.Second)
}
