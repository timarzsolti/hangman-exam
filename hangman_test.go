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
