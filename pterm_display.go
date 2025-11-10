package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/pterm/pterm"
)

type term_disp struct {
	g        *Game
	area     *pterm.AreaPrinter
	baseText string
	input    string
	message  string
}

func NewTermDisplay(game *Game) *term_disp {
	liveText, _ := pterm.DefaultArea.Start()
	term := new(term_disp)
	term.g = game
	term.area = liveText
	term.baseText = ""
	term.input = ""
	term.message = ""
	return term
}

func (term *term_disp) Close() {
	term.area.Stop()
}

func (t *term_disp) refresh() {
	t.input = ""
	t.message = ""
	letters := runesToCSV(t.g.usedLetters)

	t.baseText = fmt.Sprintf(`Word: %s
%s
Used Letters: %s
`,
		pterm.FgLightCyan.Sprint(string(t.g.display)),
		pterm.FgRed.Sprintf("Wrong guesses: (%d/%d)", t.g.wrongGuesses, t.g.maxWrong),
		pterm.FgLightBlue.Sprint(letters),
	)

	t.input = "Guess a letter: "
	t.write()
}

func (t *term_disp) write() {
	t.area.Update(t.baseText + t.message + t.input)
}

func (t *term_disp) readLetter() rune {
	reader := bufio.NewReader(os.Stdin)
	ch, _, _ := reader.ReadRune()
	undoTheNewLine()
	if !validateInput(ch) {
		undoTheNewLine()
		t.message = pterm.Red("NOT a VALID CHARACTER!\n")
		t.write()
		return t.readLetter()
	}
	if t.g.hasBeenUsed(ch) {
		t.message = pterm.Red("ALREADY USED!\n")
		t.write()
		return t.readLetter()
	}
	return unicode.ToLower(ch)
}

func undoTheNewLine() {
	fmt.Printf("\033[1A")
}

func runesToCSV(runes []rune) string {
	if len(runes) == 0 {
		return "-"
	}
	parts := make([]string, len(runes))
	for i, r := range runes {
		parts[i] = string(r)
	}
	return strings.Join(parts, ", ")
}

func (t *term_disp) endMessage() {
	if t.g.isLost() {
		t.message = pterm.FgRed.Sprintf("You lost! Your guessed wrong %d times. The word was: %s\n", t.g.wrongGuesses, t.g.word)
		t.input = ""
		t.write()
	}
	if t.g.isWon() {
		t.message = pterm.Green("Congratulations! You won!")
		t.input = ""
		t.write()
	}
}
