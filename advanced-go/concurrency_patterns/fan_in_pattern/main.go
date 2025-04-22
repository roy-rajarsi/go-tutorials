package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generator(waitGroupLogGenerators *sync.WaitGroup, serviceName string) <- chan string {
	channel := make(chan string)
	go func(waitGroupLogGenerators *sync.WaitGroup, channel chan <- string, serviceName string) {
		for i := range(10) {
			log := fmt.Sprintf("%s :: Log_%d", serviceName, i)
			if serviceName == "Service1" {
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			} else {
				time.Sleep(time.Duration(rand.Intn(5e3)) * time.Millisecond)
			}

			channel <- log
		}
		close(channel)
		waitGroupLogGenerators.Done()
	}(waitGroupLogGenerators, channel, serviceName)
	return channel
}

func getLogChannel(waitGroupLogForwarders *sync.WaitGroup, service1Channel, service2Channel <- chan string) <- chan string {
	logChannel := make(chan string)

	// We run 2 goroutines, where we pipeline the logs from service channels to the log channel

	go func() {
		for log := range(service1Channel) {
			logChannel <- log
		}
		waitGroupLogForwarders.Done()
	}()

	go func() {
		for log := range(service2Channel) {
			logChannel <- log
		}
		waitGroupLogForwarders.Done()
	}()

	go func() {
		waitGroupLogForwarders.Wait()
		close(logChannel)
	}()
	
	return logChannel
}

func main() {
	waitGroupLogGenerators := &sync.WaitGroup{}
	waitGroupLogGenerators.Add(2)
	service1Channel := generator(waitGroupLogGenerators, "Service1")
	service2Channel := generator(waitGroupLogGenerators, "Service2")
	
	waitGroupLogForwarders := &sync.WaitGroup{}
	waitGroupLogForwarders.Add(2)
	logChannel := getLogChannel(waitGroupLogForwarders, service1Channel, service2Channel)

	for log := range(logChannel) {
		fmt.Println(log)
	}
	waitGroupLogGenerators.Wait()
	waitGroupLogForwarders.Wait()
}
