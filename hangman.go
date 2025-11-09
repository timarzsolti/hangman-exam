package main

import "unicode"

type Game struct {
	word         string
	wrongGuesses int
	usedLetters  map[byte]bool
	display      []byte
	maxWrong     int
}

func NewGame(word string) *Game {
	display := make([]byte, len(word))
	for i := range display {
		display[i] = '_'
	}

	return &Game{
		word:         word,
		display:      display,
		wrongGuesses: 0,
		maxWrong:     10,
		usedLetters:  make(map[byte]bool),
	}
}

func (g *Game) GuessLetter(letter byte) bool {
	letter = toLower(letter)
	if g.usedLetter(letter) {
		return false
	}
	g.usedLetters[letter] = true

	correct := false
	for i, c := range g.word {
		ch := byte(c)
		if ch == letter {
			g.display[i] = ch
			correct = true
		}
	}

	if !correct {
		g.wrongGuesses++
	}
	return true
}

func (g *Game) usedLetter(letter byte) bool {
	return g.usedLetters[letter]
}

func toLower(c byte) byte {
	return byte(unicode.ToLower(rune(c)))
}
