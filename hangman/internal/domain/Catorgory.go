package domain

type Category string

const (
	Food      Category = "Eda"
	Animals   Category = "Животные"
	Clothes   Category = "Одежда"
	Countries Category = "Страны"
)

func (c Category) String() string {
	return string(c)
}

func (c Category) GetValues() []WordParameter {
	return []WordParameter{Food, Animals, Clothes, Countries}
}
