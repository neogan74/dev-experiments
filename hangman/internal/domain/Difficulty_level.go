package domain

type DifficultyLevel struct {
	MistakesNumber int
	Name           string
	Multiplier     float32
}

var (
	Hard   = DifficultyLevel{MistakesNumber: 10, Name: "Сложный", Multiplier: 1.0}
	Medium = DifficultyLevel{MistakesNumber: 15, Name: "Средний", Multiplier: 1.5}
	Easy   = DifficultyLevel{MistakesNumber: 18, Name: "Легкий", Multiplier: 2.0}
)

func (d DifficultyLevel) String() string {
	return d.Name
}

func (d DifficultyLevel) GetValues() []WordParameter {
	return []WordParameter{Easy, Medium, Hard}
}
