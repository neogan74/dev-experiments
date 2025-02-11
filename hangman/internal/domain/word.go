package domain

type Word struct {
	Category        Category
	DifficultyLevel DifficultyLevel
	Text            string
	Hints           []string
}

func NewWord(category Category, level DifficultyLevel, text string, hints ...string) *Word {
	return &Word{
		Category:        category,
		DifficultyLevel: level,
		Text:            text,
		Hints:           hints,
	}
}
