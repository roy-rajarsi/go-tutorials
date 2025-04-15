package main

import (
	"errors"
	"fmt"
	"time"
)

func goroutine() error {
	fmt.Println("Running Goroutine")
	time.Sleep(1 * time.Second)
	return errors.New("Error returned from Go Routine")
}

func main() {
	var err error
	fmt.Println("Running Main")
	
	go func() {
		err = goroutine()
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(err)
}
