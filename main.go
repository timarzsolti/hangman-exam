package main

import (
	"fmt"
)

func main() {
	ws := NewWordlSelector("words.txt")
	selectedWord := ws.SelectRandomWord()

	fmt.Println(selectedWord)
}
