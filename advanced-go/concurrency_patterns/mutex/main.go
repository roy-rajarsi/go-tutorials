package main

import (
	"fmt"
	"sync"
)

type Counter struct{
	Value int
	mutex sync.Mutex
}

func (counter *Counter) Increment(value int) {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	fmt.Printf("Inrementing Counter :: %d + %d -> %d\n", counter.Value, value, counter.Value+value)
	counter.Value += value
}

func (counter *Counter) Decrement(value int) {
	counter.mutex.Lock()
	counter.mutex.TryLock()
	defer counter.mutex.Unlock()
	fmt.Printf("Derementing Counter :: %d - %d -> %d\n", counter.Value, value, counter.Value-value)
	counter.Value -= value
}

func (counter *Counter) getCounterValue() (Value int) {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	return counter.Value
}

func main() {
	counter := Counter{
		Value: 100,
		mutex: sync.Mutex{},
	}

	numWorkers := 2
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(numWorkers)
	
	go func(workerId int) {
		defer waitGroup.Done()
		for range(10) {
			counter.Increment(10)
			fmt.Printf("Worker_%d :: Incremented Counter by %d\n", workerId, 10)
		}
	}(1)

	go func(workerId int) {
		defer waitGroup.Done()
		for range(5) {
			counter.Decrement(10)
			fmt.Printf("Worker_%d :: Decremented Counter by %d\n", workerId, 10)
		}
	}(2)

	waitGroup.Wait()
	fmt.Printf("Counter: %d\n", counter.getCounterValue())
}