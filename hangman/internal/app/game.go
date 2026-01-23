package app

import (
	"strings"
	"unicode"

	"github.com/neogan74/dev-experiments/hangman/internal/domain"
)

type GameStatus int

const (
	StatusInProgress GameStatus = iota
	StatusWon
	StatusLost
)

type Game struct {
	TargetWord        *domain.Word
	GuessedLetters    map[rune]bool
	RemainingAttempts int
	Status            GameStatus
}

func NewGame(word *domain.Word) *Game {
	return &Game{
		TargetWord:        word,
		GuessedLetters:    make(map[rune]bool),
		RemainingAttempts: word.DifficultyLevel.MistakesNumber,
		Status:            StatusInProgress,
	}
}

func (g *Game) Guess(letter rune) {
	if g.Status != StatusInProgress {
		return
	}

	lowerLetter := unicode.ToLower(letter)
	if g.GuessedLetters[lowerLetter] {
		return
	}

	g.GuessedLetters[lowerLetter] = true

	// Check if the guessed letter is in the target word
	found := false
	for _, char := range g.TargetWord.Text {
		if unicode.ToLower(char) == lowerLetter {
			found = true
			break
		}
	}

	if !found {
		g.RemainingAttempts--
	}

	g.updateStatus()
}

func (g *Game) updateStatus() {
	if g.RemainingAttempts <= 0 {
		g.Status = StatusLost
		return
	}

	// Check if all letters have been guessed
	allGuessed := true
	for _, char := range g.TargetWord.Text {
		if !g.GuessedLetters[unicode.ToLower(char)] {
			allGuessed = false
			break
		}
	}

	if allGuessed {
		g.Status = StatusWon
	}
}

func (g *Game) GetHiddenWord() string {
	var result strings.Builder
	for _, char := range g.TargetWord.Text {
		if g.GuessedLetters[unicode.ToLower(char)] {
			result.WriteRune(char)
		} else {
			result.WriteRune('_')
		}
		result.WriteRune(' ')
	}
	return strings.TrimSpace(result.String())
}
