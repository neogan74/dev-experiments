package app

import (
	"math/rand"
	"time"

	"github.com/neogan74/dev-experiments/hangman/internal/domain"
)

type WordRepository struct {
	words []*domain.Word
}

func NewWordRepository() *WordRepository {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	return &WordRepository{
		words: []*domain.Word{
			domain.NewWord(domain.Food, domain.Easy, "Apple", "A red fruit"),
			domain.NewWord(domain.Food, domain.Medium, "Pizza", "Italian dish"),
			domain.NewWord(domain.Animals, domain.Easy, "Cat", "Meow"),
			domain.NewWord(domain.Animals, domain.Hard, "Elephant", "Big ears, long trunk"),
			domain.NewWord(domain.Countries, domain.Medium, "France", "Paris is the capital"),
			domain.NewWord(domain.Clothes, domain.Easy, "Shirt", "Upper body wear"),
		},
	}
}

func (r *WordRepository) GetRandomWord() *domain.Word {
	if len(r.words) == 0 {
		return nil
	}
	return r.words[rand.Intn(len(r.words))]
}
