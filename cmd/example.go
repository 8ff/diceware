package main

import (
	"diceware"
	"fmt"
)

func main() {
	// Get a slice of words
	words := diceware.GetWords()
	fmt.Println("Number of entries:", len(words))

	// Get a randomised slice of words
	randomWords := diceware.GetRandomWords()
	fmt.Println("Number of entries in randomised list:", len(randomWords))

	// Get a map of words
	wordsMap := diceware.GetWordsMap()
	fmt.Println("Number of entries in words map:", len(wordsMap))
}
