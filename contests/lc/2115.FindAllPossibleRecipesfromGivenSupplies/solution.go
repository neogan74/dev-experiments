package main

import (
	"fmt"
)

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	graph := make(map[string][]string)
	indegree := make(map[string]int)
	available := make(map[string]bool)

	// Начальные ингредиенты доступны
	for _, s := range supplies {
		available[s] = true
	}

	// Построение графа зависимостей
	for i, recipe := range recipes {
		for _, ing := range ingredients[i] {
			graph[ing] = append(graph[ing], recipe)
			indegree[recipe]++
		}
	}

	queue := []string{}
	// В очередь кладём изначально доступные ингредиенты
	for s := range available {
		queue = append(queue, s)
	}

	result := []string{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Если текущий узел — рецепт, добавляем в результат
		for _, next := range graph[curr] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
				available[next] = true
				result = append(result, next)
			}
		}
	}

	return result
}

func main() {
	fmt.Println(findAllRecipes(
		[]string{"bread"},
		[][]string{{"yeast", "flour"}},
		[]string{"yeast", "flour", "corn"}))
	// ["bread"]

	fmt.Println(findAllRecipes(
		[]string{"bread", "sandwich"},
		[][]string{{"yeast", "flour"}, {"bread", "meat"}},
		[]string{"yeast", "flour", "meat"}))
	// ["bread", "sandwich"]

	fmt.Println(findAllRecipes(
		[]string{"bread", "sandwich", "burger"},
		[][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}},
		[]string{"yeast", "flour", "meat"}))
	// ["bread", "sandwich", "burger"]
}
