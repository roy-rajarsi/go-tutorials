package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {
	connection string
}

var database *Database
var databaseConenctionInitiator sync.Once

func getDatabase() *Database {
	databaseConenctionInitiator.Do(
		func() {  // This function is run only once
			fmt.Println("Creating Database Connection")
			database = &Database{connection: "ConnectionString1234"}
			time.Sleep(2 * time.Second)
			fmt.Println("Database Connection Created")
		})
	return database
}

func main() {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(5)
	for range(5) {
		go func(waitGroup *sync.WaitGroup) {
			defer waitGroup.Done()
			fmt.Println("Database instance fetched :: ", getDatabase()) // getDatabase() is called by all goroutines
		}(waitGroup)
	}
	waitGroup.Wait()
}