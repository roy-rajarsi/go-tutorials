package main

import (
	"fmt"
	"time"
)

func poll(jobStartedAt time.Time) {
	fmt.Println("Polling started at: ", jobStartedAt)
	time.Sleep(2 * time.Second)
	fmt.Println("Polling ended at ", time.Now())
	fmt.Println()
}

func main() {
	interval := time.Duration(5 * time.Second)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for tick := range(ticker.C) {
		go poll(tick)
	}
}