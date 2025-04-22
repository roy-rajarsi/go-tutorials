package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func reader(waitGroup *sync.WaitGroup, mutex *sync.RWMutex, criticalSection *int) {
	defer waitGroup.Done() 
	defer mutex.RUnlock()

	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	mutex.RLock()
	fmt.Println("Read CS :: ", *criticalSection)
}

func writer(waitGroup *sync.WaitGroup, mutex *sync.RWMutex, criticalSection *int) {
	defer waitGroup.Done()
	defer mutex.Unlock()

	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	mutex.Lock()
	fmt.Println("Wrting CS :: 200 from ", *criticalSection)
	*criticalSection = 200
}

func main() {
	waitGroup := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}
	criticalSection := 10
	waitGroup.Add(2)
	go reader(waitGroup, mutex, &criticalSection)
	go writer(waitGroup, mutex, &criticalSection)

	waitGroup.Wait()
}