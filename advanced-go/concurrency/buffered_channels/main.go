package main

import "fmt"

func main() {
	bufferedChannel := make(chan string, 5)  // Creates a Buffered String Channel of Size 5
	bufferedChannel <- "Hello"
	bufferedChannel <- "I"
	bufferedChannel <- "am"
	bufferedChannel <- "learning"
	bufferedChannel <- "Go"

	// Now, the buffer in the Channel is full. Someone must start listening to this channel, 
	// before any write is performed on this channel

	bufferedChannel <- "This String Will cause Deadlock !"
	fmt.Println(bufferedChannel)
}
