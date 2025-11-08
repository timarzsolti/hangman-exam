package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var words []string

func main() {
	var err error
	words, err = loadWordsFromFile("words.txt")
	if err != nil {
		setupPredefinedWords()
	}

	fmt.Println("Loaded words:", words)
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

func loadWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			words = append(words, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
