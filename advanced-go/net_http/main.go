package main

import (
	"io"
	"fmt"
	"net/http"
)


func startClient() {
	httpClient := &http.Client{}
	response, err := httpClient.Get("https://example.com")

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body:: ", string(body))
	fmt.Println("Response Status:: ", response.Status)
}

func main() {
	startClient()
}
