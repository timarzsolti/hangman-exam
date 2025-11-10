package main

import (
	"fmt"
)

func main() {
	ws := NewWordlSelector("words.txt")
	selectedWord := ws.SelectRandomWord()

	hman := NewGame(selectedWord)
	terminal := NewTermDisplay(hman)
	defer terminal.Close()
	for !hman.isOver() {
		terminal.refresh()
		ch := terminal.readLetter()
		if !hman.GuessLetter(ch) {
			//write error for user
		}
	}
	terminal.refresh()
	terminal.endMessage()
}

func writeUsedLetterError() {
	fmt.Println("The letter was already used")
}
