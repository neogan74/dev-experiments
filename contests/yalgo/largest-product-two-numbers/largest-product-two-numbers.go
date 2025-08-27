package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var numbers []int64
	for scanner.Scan() {
		var num int64
		fmt.Sscan(scanner.Text(), &num)
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if len(numbers) < 2 {
		fmt.Println("List must contain at least two numbers.")
		return
	}

	maxPositive1, maxPositive2 := int64(math.MinInt64), int64(math.MinInt64)
	minNegative1, minNegative2 := int64(math.MaxInt64), int64(math.MaxInt64)

	for _, num := range numbers {
		if num > maxPositive1 {
			maxPositive2 = maxPositive1
			maxPositive1 = num
		} else if num > maxPositive2 {
			maxPositive2 = num
		}

		if num < minNegative1 {
			minNegative2 = minNegative1
			minNegative1 = num
		} else if num < minNegative2 {
			minNegative2 = num
		}
	}

	if maxPositive1*maxPositive2 > minNegative1*minNegative2 {
		if maxPositive1 > maxPositive2 {
			fmt.Println(maxPositive2, maxPositive1)
		} else {
			fmt.Println(maxPositive1, maxPositive2)
		}
	} else {
		if minNegative1 < minNegative2 {
			fmt.Println(minNegative1, minNegative2)
		} else {
			fmt.Println(minNegative2, minNegative1)
		}
	}
}
