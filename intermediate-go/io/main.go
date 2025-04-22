package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	source := "Hello, Go I/O!"
	stringReader := strings.NewReader(source)  // strings.NewReader() returns a *strings.Reader, which implements io.Reader

	buffer := make([] byte, 4)
	for {
		bytesRead, error := stringReader.Read(buffer)  // Read a maximum of len(buffer) from source to buffer
		if error == io.EOF {  // If while reading these 4 bytes, the file ends, we receive an io.EOF error
			fmt.Println("Reading Completed ...")
			break
		}
		fmt.Printf("Read %d bytes.\nBytes Read: %s\n\n", bytesRead, buffer[:bytesRead])
	}

	stringToWrite := "Hello Golang Writer !\n"
	standardOutput := os.Stdout  // Returns *os.File, which implements io.Writer
	bytesWritten, err := standardOutput.Write([]byte(stringToWrite))  // Convert the string to a byte-slice and write it entirely

	if err != nil {
		fmt.Println("Error while writing: ", err)
		return
	}
	fmt.Println("Bytes Written: ", bytesWritten)
}
