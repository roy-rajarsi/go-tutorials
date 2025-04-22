package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func post(ctx context.Context, waitGroup *sync.WaitGroup, id int, url string) {
	requestTime := time.Duration(rand.Intn(5e3)) * time.Millisecond
	fmt.Println("Goroutine_", id, ":: Launched and will take:", requestTime)
	for {
		select {
		case <- ctx.Done():
			fmt.Println("Goroutine_", id, ":: Post Request on", url, "for requestId:", ctx.Value("requestId"), "is cancelled")
			fmt.Println(ctx.Err())
			waitGroup.Done()
			return
		case <- time.After(requestTime):
			fmt.Println("Goroutine_", id, ":: Post Request on", url, "for requestId:", ctx.Value("requestId"), "is completed")
			waitGroup.Done()
			return
		}
	}
}

func main() {
	requestId := "2341"
	url := "localhost"
	requestExpiry := 2 * time.Second
	baseContext := context.Background()

	// Associate some request-scoped values in the context, and create a new context
	contextWithRequestParams := context.WithValue(baseContext, "requestId", requestId)

	// func context.WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
	// The done channel of the context returned will be closed automatically after "timeout"
	// If for some reason, we wish to close the Done channel before this timeout, we can simply invoke the cancel() returned
	contextWithTimeout, cancelFunction := context.WithTimeout(contextWithRequestParams, time.Duration(requestExpiry))
	defer cancelFunction()

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(3)

	go post(contextWithTimeout, waitGroup, 1, url)
	go post(contextWithTimeout, waitGroup, 2, url)
	go post(contextWithTimeout, waitGroup, 3, url)

	waitGroup.Wait()

}
