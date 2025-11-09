package main

import "unicode"

type Game struct {
	word         string
	wrongGuesses int
	usedLetters  map[byte]bool
	display      []byte
	charsGuessed int
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
		charsGuessed: 0,
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
			g.charsGuessed++
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

func (g *Game) isWon() bool {
	return (g.charsGuessed >= len(g.word))
}

func (g *Game) isLost() bool {
	return g.wrongGuesses >= g.maxWrong
}
