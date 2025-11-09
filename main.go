package main

import (
	"fmt"
)

func main() {
	ws := NewWordlSelector("words.txt")
	selectedWord := ws.SelectRandomWord()

	hman := NewGame(selectedWord)
	for shouldContinue(hman) {
		writeToConsole(hman)
		character := readLetter()
		ch := byte(character)
		if !hman.GuessLetter(ch) {
			writeUsedLetterError()
		}
	}
	writeToConsole(hman)
	if hman.isLost() {
		fmt.Printf("You lost! Your guessed wrong %d times. The word was: %s\n", hman.wrongGuesses, hman.word)
	}
	if hman.isWon() {
		fmt.Printf("Congratulations! You won!\n")
	}
}

func writeUsedLetterError() {
	fmt.Println("The letter was already used")
}

func shouldContinue(g *Game) bool {
	return !g.isLost() && !g.isWon()
}
