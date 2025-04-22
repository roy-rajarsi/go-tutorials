package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(channel chan <- string) {
	for {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		message := fmt.Sprintf("Message%d", rand.Int())
		channel <- message
	}
}

func consumer(channel <- chan string) {
	for message := range(channel) {
		fmt.Println("Consumer :: Message received -> ", message)
	}
}

func main() {
	channel := make(chan string)
	go producer(channel)
	go consumer(channel)
	time.Sleep(10 * time.Second)
}
