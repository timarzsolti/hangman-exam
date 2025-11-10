package main

import (
	"slices"
	"unicode"
)

type Game struct {
	word         string
	wrongGuesses int
	usedLetters  []rune
	display      []rune
	charsGuessed int
	maxWrong     int
}

func NewGame(word string) *Game {
	display := make([]rune, len(word))
	for i := range display {
		display[i] = '_'
	}

	return &Game{
		word:         word,
		display:      display,
		wrongGuesses: 0,
		maxWrong:     10,
		usedLetters:  make([]rune, 0),
		charsGuessed: 0,
	}
}

func (g *Game) usingALetter(letter rune) {
	g.usedLetters = append(g.usedLetters, letter)
}

func (g *Game) GuessLetter(letter rune) bool {
	letter = unicode.ToLower(letter)
	if g.hasBeenUsed(letter) {
		return false
	}
	g.usingALetter(letter)

	correct := false
	for i, ch := range g.word {
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

func (g *Game) hasBeenUsed(letter rune) bool {
	return slices.Contains(g.usedLetters, letter)
}

func (g *Game) isWon() bool {
	return (g.charsGuessed >= len(g.word))
}

func (g *Game) isLost() bool {
	return g.wrongGuesses >= g.maxWrong
}

func (g *Game) isOver() bool {
	return g.isLost() || g.isWon()
}
