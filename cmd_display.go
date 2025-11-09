package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func writeToConsole(game *Game) {
	fmt.Print("Word: ")
	for _, ch := range game.display {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	fmt.Printf("Wrong guesses: %d\n", game.wrongGuesses)
}

func readLetter() rune {
	fmt.Print("Guess a letter: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	ch := getFirstRune(input)
	if !validateInput(ch) {
		return readLetter()
	}
	return unicode.ToLower(ch)
}

func getFirstRune(s string) rune {
	for _, letter1 := range s {
		return letter1
	}
	return 0
}

func validateInput(ch rune) bool {
	if !unicode.IsLetter(ch) {
		fmt.Println("Invalid Letter")
		return false
	}
	return true
}
