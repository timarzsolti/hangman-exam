package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuessLetter(t *testing.T) {
	g := NewGame("apple")

	goodChar := g.GuessLetter('p')
	assert.True(t, goodChar, "GuessLetter should return true if the character was not entered yet")
	assert.Equal(t, []byte{'_', 'p', 'p', '_', '_'}, g.display, "Display should update correctly")

	goodChar = g.GuessLetter('p')
	assert.False(t, goodChar, "GuessLetter should return false if the character was already entered")
	assert.Equal(t, []byte{'_', 'p', 'p', '_', '_'}, g.display, "display should not change")
	assert.Equal(t, 0, g.wrongGuesses, "internal state should not change")

	goodChar = g.GuessLetter('z')
	assert.True(t, goodChar, "GuessLetter should return true if the character was not entered yet")
	assert.Equal(t, 1, g.wrongGuesses, "WrongGuesses should increment")

	goodChar = g.GuessLetter('z')
	assert.False(t, goodChar, "GuessLetter should return false if the character was already entered")
	assert.Equal(t, 1, g.wrongGuesses, "WrongGuesses should not change if the character was already entered")

}

func TestGuessLetter_UpperCaseDoesNotMatter(t *testing.T) {
	g := NewGame("apple")

	available := g.GuessLetter('z')
	assert.True(t, available, "GuessLetter should return true if the character was not entered yet")
	assert.Equal(t, 1, g.wrongGuesses, "WrongGuesses should increment")

	available = g.GuessLetter('Z')
	assert.False(t, available, "Should be considered equal with the lower one")
	assert.Equal(t, 1, g.wrongGuesses, "internal state should not change")
}

func TestIsWon(t *testing.T) {
	g := NewGame("word")
	assert.False(t, g.isWon(), "Did not reveal all the letters")
	g.GuessLetter('w')
	assert.False(t, g.isWon(), "Did not reveal all the letters")
	g.GuessLetter('o')
	assert.False(t, g.isWon(), "Did not reveal all the letters")
	g.GuessLetter('d')
	assert.False(t, g.isWon(), "Did not reveal all the letters")
	g.GuessLetter('r')
	assert.True(t, g.isWon(), "All the letters are revealed")
}

func TestIsLost(t *testing.T) {
	g := NewGame("word")
	g.maxWrong = 2
	assert.Equal(t, 0, g.wrongGuesses, "No guesses made yet")
	assert.False(t, g.isLost(), "Did not reach the max missed letters")
	g.GuessLetter('w')
	assert.Equal(t, 0, g.wrongGuesses, "it was a good guess")
	assert.False(t, g.isLost(), "Did not reach the max missed letters")
	g.GuessLetter('x')
	assert.Equal(t, 1, g.wrongGuesses, "it was a wrong guess")
	assert.False(t, g.isLost(), "Did not reach the max missed letters")
	g.GuessLetter('d')
	assert.Equal(t, 1, g.wrongGuesses, "it was a good guess")
	assert.False(t, g.isLost(), "Did not reach the max missed letters")
	g.GuessLetter('y')
	assert.Equal(t, 2, g.wrongGuesses, "it was a wrong guess")
	assert.True(t, g.isLost(), "Did reach the max missed letters")
}
