package main

import (
	"fmt"
	"strings"
)

func main() {
	var message string = "Hello, World!"
	fmt.Printf("Input String = %s", message)

	fmt.Println() // Introduce new line
	// Comments in Go
	for index, character := range message {
		fmt.Printf("\tPosition: %d, Character: %c", index, character)
		fmt.Println()
	}

	var inputStream string = "a basic set of words to iterate upon until completion"
	inputStreamArr := strings.Split(inputStream, " ")
	for idx := 0; idx < len(inputStreamArr); idx++ {
		fmt.Printf("\tCurrent Word = %s", inputStreamArr[idx])
		fmt.Println()
	}

}
