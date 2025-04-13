package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := "Hello I am learning about the Bufio Package in Go !\nIt is great.\n"
	stringReader := strings.NewReader(data)  // Returns a *strings.Reader
	bufferedReader := bufio.NewReader(stringReader)  // Returns a *bufio.Reader

	// Reading Upto N bytes
	countOfCharactersToRead := 20
	charactersRead := make([]byte, countOfCharactersToRead)
	countOfCharactersRead, error := bufferedReader.Read(charactersRead)  // Read N bytes and store it into charactersRead

	if error != nil {
		fmt.Println("Error occured while reading: ", error)
		return
	}

	fmt.Printf("Count of Characters Read: %d\nCharacters Read:\n%s\n", countOfCharactersRead, charactersRead)

	// Read Lines
	stringReader2 := strings.NewReader(data)  // Returns a *strings.Reader
	bufferedReader2 := bufio.NewReader(stringReader2)  // Returns a *bufio.Reader
	firstLine, error := bufferedReader2.ReadString('\n')
	if error != nil {
		fmt.Println("Error occured while reading: ", error)
		return
	}

	fmt.Printf("First Line: %s", firstLine)

	secondLine, error := bufferedReader2.ReadString('\n')
	if error != nil {
		fmt.Println("Error occured while reading: ", error)
		return
	}

	fmt.Printf("Second Line: %s", secondLine)

	stdOutWriter := os.Stdout
	bufferedWriter := bufio.NewWriter(stdOutWriter)
	
	toWriteByteSlice := []byte("This Byte Slice Needs To Be Written")
	charactersWritten, error := bufferedWriter.Write(toWriteByteSlice)
	if error != nil {
		fmt.Println("Error while writing: ", error)
		return
	}
	bufferFlushError := bufferedWriter.Flush()
	if bufferFlushError != nil {
		fmt.Println("Error while flushing: ", bufferFlushError)
		return
	}
	fmt.Printf("\nBytes Written: %d\n", charactersWritten)

	toWriteString := "This String Needs To Be Written"
	stdOutWriter2 := os.Stdout
	bufferedWriter2 := bufio.NewWriter(stdOutWriter2)
	charactersWritten2, error2 := bufferedWriter2.WriteString(toWriteString)
	if error2 != nil {
		fmt.Println("Error while writing: ", error2)
		return
	}
	bufferFlushError2 := bufferedWriter2.Flush()
	if bufferFlushError2 != nil {
		fmt.Println("Error while flushing: ", bufferFlushError2)
		return
	}
	fmt.Printf("\nBytes Written: %d\n", charactersWritten2)
}
