package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// TODO: can be refactored for using StdOut with flags.
	part1()
	part2()
}

func part1() {
	// Стартовая позиция
	position := 50
	// Счётчик остановок на нуле
	zeroCount := 0

	// Открываем файл с входными данными
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}

		direction := line[0]
		valueStr := line[1:]
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println("Ошибка парсинга числа:", err)
			continue
		}

		// Обновляем позицию
		if direction == 'R' {
			position = (position + value) % 100
		} else if direction == 'L' {
			position = (position - value + 100) % 100
		}

		// Проверяем, остановились ли мы на нуле
		if position == 0 {
			zeroCount++
		}
	}

	// Выводим результат
	fmt.Printf("Количество остановок на нуле: %d\n", zeroCount)
}

func part2() {
	position := 50
	zeroCount := 0

	// Открываем файл с входными данными
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}

		direction := line[0]
		valueStr := line[1:]
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println("Ошибка парсинга числа:", err)
			continue
		}

		for step := 1; step <= value; step++ {
			if direction == 'R' {
				position = (position + 1) % 100
			} else if direction == 'L' {
				position = (position - 1 + 100) % 100
			}

			if position == 0 {
				zeroCount++
			}
		}
	}

	fmt.Printf("Количество остановок на нуле: %d\n", zeroCount)
}
