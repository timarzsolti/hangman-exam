package main

import (
	"fmt"
)

func main() {
	ws := WordSelector{}
	ws.SetupWords("words.txt")
	selectedWord := ws.SelectRandomWord()

	fmt.Println(selectedWord)
}
