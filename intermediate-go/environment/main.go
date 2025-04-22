package main

import (
	"fmt"
	"os"
)

func main() {
	user := os.Getenv("USER")  // rajarsiroy
	home := os.Getenv("HOME")  // Users/rajarsiroy
	
	fmt.Println("USER ::", user)
	fmt.Println("HOME ::", home)
	
	fmt.Println("\nAll Environment Variables ...\n\n")
	allEnvironmentVariables := os.Environ()
	for index, enviromentVariable := range(allEnvironmentVariables) {
	  fmt.Println(index, enviromentVariable)
	}

	os.Setenv("MyEnvironmentVariable", "Hello From Go !")
	fmt.Println("\n\nMyEnvironmentVariable :: ", os.Getenv("MyEnvironmentVariable")) // Hello From Go !
	
	os.Unsetenv("MyEnvironmentVariable")
	fmt.Println("MyEnvironmentVariable (Unset) :: ", os.Getenv("MyEnvironmentVariable"))  // ""
}