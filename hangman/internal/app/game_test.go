package app

import (
	"testing"

	"github.com/neogan74/dev-experiments/hangman/internal/domain"
)

func TestNewGame(t *testing.T) {
	word := domain.NewWord(domain.Food, domain.Easy, "Test", "Hint")
	game := NewGame(word)

	if game.TargetWord != word {
		t.Errorf("Expected TargetWord to be %v, got %v", word, game.TargetWord)
	}
	if game.RemainingAttempts != word.DifficultyLevel.MistakesNumber {
		t.Errorf("Expected RemainingAttempts to be %d, got %d", word.DifficultyLevel.MistakesNumber, game.RemainingAttempts)
	}
	if game.Status != StatusInProgress {
		t.Errorf("Expected Status to be InProgress, got %v", game.Status)
	}
}

func TestGuess(t *testing.T) {
	word := domain.NewWord(domain.Food, domain.Easy, "Test", "Hint")
	game := NewGame(word)

	// Correct guess
	game.Guess('t')
	if !game.GuessedLetters['t'] {
		t.Error("Expected 't' to be guessed")
	}
	if game.RemainingAttempts != word.DifficultyLevel.MistakesNumber {
		t.Errorf("Expected RemainingAttempts not to decrease, got %d", game.RemainingAttempts)
	}

	// Incorrect guess
	game.Guess('z')
	if !game.GuessedLetters['z'] {
		t.Error("Expected 'z' to be guessed")
	}
	if game.RemainingAttempts != word.DifficultyLevel.MistakesNumber-1 {
		t.Errorf("Expected RemainingAttempts to decrease, got %d", game.RemainingAttempts)
	}

	// Repeated guess
	game.Guess('t')
	if game.RemainingAttempts != word.DifficultyLevel.MistakesNumber-1 {
		t.Errorf("Expected RemainingAttempts not to decrease on repeated guess, got %d", game.RemainingAttempts)
	}
}

func TestGameWin(t *testing.T) {
	word := domain.NewWord(domain.Food, domain.Easy, "Go", "Hint")
	game := NewGame(word)

	game.Guess('g')
	game.Guess('o')

	if game.Status != StatusWon {
		t.Errorf("Expected Status to be Won, got %v", game.Status)
	}
}

func TestGameLoss(t *testing.T) {
	word := domain.NewWord(domain.Food, domain.Easy, "Hi", "Hint")
	// Set mistakes to 1 for easier testing, but since it's struct field and not modifiable easily here without changing domain,
	// I can just exhaust the attempts.
	// Actually domain.Easy has 18 mistakes. That's a lot for a test loop.
	// I will mock a word with a difficulty level manually if possible, or just loop.

	// We can just set the game.RemainingAttempts manually for testing purpose since it's exported
	game := NewGame(word)
	game.RemainingAttempts = 1

	game.Guess('z')

	if game.Status != StatusLost {
		t.Errorf("Expected Status to be Lost, got %v", game.Status)
	}
}
