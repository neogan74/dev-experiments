package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	position := 50
	zeroCount := 0

	scanner := bufio.NewScanner(os.Stdin)

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
