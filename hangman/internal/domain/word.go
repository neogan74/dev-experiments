package domain

type Word struct {
	Category        Category
	DifficultyLevel DifficultyLevel
	Text            string
	Hints           []string
}

func NewWord(category Category)
