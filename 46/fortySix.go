package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"io"
	"time"
)

func main() {
	filename := getNonEmptyInput("Name of file to process: ")
	start := time.Now()
	ok, wordFrequencyMap, frequencyToWordsMap := processFile(filename)
	if ok {
		displayHistogram(wordFrequencyMap, frequencyToWordsMap)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("The execution took %.2f seconds.\n", secs)
}

func getNonEmptyInput(msg string) (input string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			done = true
		}
	}

	return input
}


func processFile(filename string) (ok bool, wordFrequencyMap map[string]int, frequencyToWordsMap map[int]map[string]bool){
	wordFrequencyMap = make(map[string]int, 0)
	frequencyToWordsMap = make(map[int]map[string]bool, 0)

	file, e := os.Open(filename)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not open file %s!\n", filename)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	done := false
	for !done {
		line, e := reader.ReadString('\n')
		if e != nil {
			done = true
		}

		if e == nil || e == io.EOF {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.TrimSpace(word)
				if len(word) == 0 { // Disregard ""
					continue
				}
				delete(frequencyToWordsMap[wordFrequencyMap[word]], word)
				if len(frequencyToWordsMap[wordFrequencyMap[word]]) == 0 {
					delete(frequencyToWordsMap, wordFrequencyMap[word])
				}

				wordFrequencyMap[word]++

				if _, elementInMap := frequencyToWordsMap[wordFrequencyMap[word]]; !elementInMap  {
					frequencyToWordsMap[wordFrequencyMap[word]] = make(map[string]bool)
				}
				frequencyToWordsMap[wordFrequencyMap[word]][word] = true
			}
		}
	}

	ok = true
	return
}

func displayHistogram(wordFrequencyMap map[string]int, frequencyToWordsMap map[int]map[string]bool) {
	frequenciesInSortedOrder := returnFrequenciesInSortedOrder(frequencyToWordsMap)
	lengthOfLongestWord := findLengthOfLongestWord(wordFrequencyMap)
	for _, frequency := range frequenciesInSortedOrder {
		for word, _ := range frequencyToWordsMap[frequency] {
			fmt.Print(word)
			fmt.Print(":")
			fmt.Print(strings.Repeat(" ", lengthOfLongestWord - len(word) + 1))
			fmt.Println(strings.Repeat("*", frequency))
		}
	}
}

func returnFrequenciesInSortedOrder(frequencyToWordsMap map[int]map[string]bool) (frequenciesInSortedOrder[]int) {
	frequenciesInSortedOrder = make([]int, len(frequencyToWordsMap))
	for frequencyToSort, _ := range frequencyToWordsMap {
		insertionPoint := -1
		placeLast := false
		for indexOfSortedFrequency, sortedFrequency := range frequenciesInSortedOrder {
			if sortedFrequency == 0 {
				insertionPoint = indexOfSortedFrequency
				placeLast = true
				break
			}

			if  sortedFrequency < frequencyToSort {
				insertionPoint = indexOfSortedFrequency
				break
			}
		}

		if !placeLast {
			copy(frequenciesInSortedOrder[insertionPoint + 1:], frequenciesInSortedOrder[insertionPoint: len(frequenciesInSortedOrder) -1])

		}
		frequenciesInSortedOrder[insertionPoint] = frequencyToSort
	}

	return
}


func findLengthOfLongestWord(wordFrequencyMap map[string]int) (lengthOfLongestWord int) {
	for word, _ := range wordFrequencyMap {
		if len(word) > lengthOfLongestWord {
			lengthOfLongestWord = len(word)
		}
	}

	return
}