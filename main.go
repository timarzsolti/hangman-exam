package main

import (
	"fmt"
	"math/rand"
)

var words []string

func main() {
	setupPredefinedWords()
	fmt.Println(selectRandomWord())
}

func setupPredefinedWords() {
	words = []string{"apple", "banana", "cherry", "date"}
}

func selectRandomWord() string {
	// as of go 1.24+ no need to run rand.Seed

	randomIndex := rand.Intn(len(words))

	return words[randomIndex]
}
