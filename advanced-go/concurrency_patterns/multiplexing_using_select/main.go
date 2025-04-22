package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	go func(channel chan string) {
		time.Sleep(4 * time.Second)
		fmt.Printf("Received %s in Channel3\n", <-channel)
	}(channel3)

	messageToWrite := "Hello World !!!!"

	select {
		case messageToRead := <- channel1: {
			fmt.Printf("Message: %s received from Channel1\n", messageToRead)
		}

		case messageToRead := <- channel2: {
			fmt.Printf("Message: %s received from Channel2\n", messageToRead)
		}

		case channel3 <- messageToWrite: {
			fmt.Printf("Message: %s written in Channel3\n", messageToWrite)
			time.Sleep(time.Second)
		}

		case <- time.After(3 * time.Second): {
			fmt.Println("No Goroutines responded... Dropping !")
		}

		// default: {
		// 	fmt.Println("Dropping !!!!")
		// }
	}
	fmt.Println("Exiting...")
}
