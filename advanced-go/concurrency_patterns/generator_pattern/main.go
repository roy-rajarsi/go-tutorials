package main

import (
	"fmt"
	"math/rand"
	"time"
)

func populateChannel(channel chan <- string) {
	for i := range(10) {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		message := fmt.Sprintf("Message_%d", i)
		channel <- message
	}
	close(channel)
}

func generator() <- chan string {
	channel := make(chan string)

	// Start a new goroutine that will populate this channel
	go populateChannel(channel)

	return channel
}

func main() {

	consumerChannel1 := generator()
	consumerChannel2 := generator()

	go func() {
		for message1 := range consumerChannel1 {
			fmt.Println("Consumer1 :: ", message1)
		}
	}()

	go func() {
		for message2 := range consumerChannel2 {
			fmt.Println("Consumer2 :: ", message2)
		}
	}()
	time.Sleep(10 * time.Second)
}
