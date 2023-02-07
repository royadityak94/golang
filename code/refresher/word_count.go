/*
	Word Count Implementation

Word count implementation using Go
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func isItemInEntries(entries []string, item string) bool {
	item = strings.TrimSpace(item)
	for _, entry := range entries {
		if entry == item {
			return true
		}
	}
	return false
}

func wordCountNaive(fileName string, topK int) map[string]int {
	// Read the input file
	filePointer, errFilePointer := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if errFilePointer != nil {
		panic("Error Opening the file: " + errFilePointer.Error())
	}
	defer filePointer.Close()

	// Words to Ignore and word counter map assignment
	wordsToIgnore := []string{"the", "and", "for", "not"}
	wordCounterMapper := make(map[string]int)

	// Reading the file line by line
	fileReader := bufio.NewScanner(filePointer)
	for fileReader.Scan() {
		line := strings.Split(strings.ToLower(strings.TrimSpace(fileReader.Text())), " ")
		for _, word := range line {
			if len(word) > 3 && (!isItemInEntries(wordsToIgnore, word)) {
				wordCounterMapper[word] += 1
			}
		}
	}
	fileReaderErr := fileReader.Err()
	if fileReaderErr != nil {
		panic("Error in processing the file: " + fileReaderErr.Error())
	}

	// Sorting the map by its value

	// Creating a list of keys
	keys := make([]string, 0, len(wordCounterMapper))
	for key := range wordCounterMapper {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return wordCounterMapper[keys[i]] > wordCounterMapper[keys[j]]
	})

	// Extract topK in another map
	resultTopK := make(map[string]int)
	cnt := 0
	for _, word := range keys {
		resultTopK[word] = wordCounterMapper[word]
		cnt += 1
		if cnt > topK {
			break
		}
	}
	return resultTopK
}

func main() {
	var fileName string = "data/word_count_sample.txt"
	var topK int = 15
	// Output from Naive
	resultTopK := wordCountNaive(fileName, topK)
	for key, value := range resultTopK {
		fmt.Printf("(%s, %d), ", key, value)
	}
	fmt.Println()
}
