package main

import (
	"math/rand"
	"fmt"
	"time"
)

func boring(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Counter: ", i, "Message: ", message)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go boring("I am a Boring Function !")
	
	fmt.Println("I am waiting for Boring() to finish...")
	time.Sleep(2 * time.Second)
	fmt.Println("I am Leaving !")
}
