package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/neogan74/dev-experiments/hangman/internal/app"
)

func main() {
	repo := app.NewWordRepository()
	word := repo.GetRandomWord()
	if word == nil {
		fmt.Println("Ошибка: Не удалось найти слова в репозитории.")
		return
	}

	game := app.NewGame(word)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Добро пожаловать в игру Виселица!")
	fmt.Printf("Категория: %s\n", word.Category)
	fmt.Printf("Сложность: %s\n", word.DifficultyLevel.Name)
	fmt.Println("Попробуйте угадать слово.")

	for game.Status == app.StatusInProgress {
		fmt.Printf("\nСлово: %s\n", game.GetHiddenWord())
		fmt.Printf("Осталось попыток: %d\n", game.RemainingAttempts)
		fmt.Print("Введите букву: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if utf8.RuneCountInString(input) != 1 {
			fmt.Println("Пожалуйста, введите одну букву.")
			continue
		}

		letter, _ := utf8.DecodeRuneInString(input)
		game.Guess(letter)
	}

	if game.Status == app.StatusWon {
		fmt.Printf("\nПоздравляем! Вы угадали слово: %s\n", word.Text)
	} else {
		fmt.Printf("\nВы проиграли. Загаданное слово было: %s\n", word.Text)
	}
}
