package main

import (
	"bufio"
	"math/rand"
	"os"
)

type WordSelector struct {
	words []string
}

func (ws *WordSelector) SetupWords(filename string) {
	words, err := ws.loadWordsFromFile(filename)
	if err != nil {
		words = ws.setupPredefinedWords()
	}
	ws.words = words
}

func (ws *WordSelector) loadWordsFromFile(filename string) ([]string, error) {
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

func (ws *WordSelector) setupPredefinedWords() []string {
	return []string{"apple", "banana", "cherry", "date"}
}

func (ws *WordSelector) SelectRandomWord() string {
	// as of go 1.24+ no need to run rand.Seed

	randomIndex := rand.Intn(len(ws.words))

	return ws.words[randomIndex]
}
