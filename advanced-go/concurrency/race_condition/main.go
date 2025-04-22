package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cs := 10

	go func(cs *int) {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		*cs = 100
	}(&cs)

	go func(cs *int) {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		*cs = 200
	}(&cs)

	time.Sleep(5 * time.Second)
	fmt.Printf("Critical Section :: %d\n", cs)
}