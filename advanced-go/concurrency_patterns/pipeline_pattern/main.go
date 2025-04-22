package main

import (
	"fmt"
	"sync"
)

func poll(waitGroup *sync.WaitGroup) <- chan int {
	messages := make(chan int)
	go func(waitGroup *sync.WaitGroup, messages chan <- int) {
		for message := range(5) {
			fmt.Printf("Poll :: Message %d\n", message)
			messages <- message
		}
		close(messages)
		waitGroup.Done()
	}(waitGroup, messages)
	return messages
}

func process(waitGroup *sync.WaitGroup, messages <- chan int) <- chan int {
	outputs := make(chan int)
	go func(waitGroup *sync.WaitGroup, messages <- chan int, outputs chan <- int) {
		for message := range(messages) {
			fmt.Printf("Process :: Received Message %d\n", message)
			fmt.Printf("Process :: Processed Message %d To %d\n", message, 2 * message)
			outputs <- 2 * message
		}
		close(outputs)
		waitGroup.Done()
	}(waitGroup, messages, outputs)
	return outputs
}

func main() {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)
	messages := poll(waitGroup)
	outputs := process(waitGroup, messages)

	for output := range(outputs) {
		fmt.Printf("Output :: %d\n", output)
	}
	waitGroup.Wait()
}
