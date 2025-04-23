package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomizedSet представляет нашу структуру данных.
type RandomizedSet struct {
	values  []int
	indices map[int]int // отображает значение -> индекс в срезе values
}

// NewRandomizedSet инициализирует новый объект RandomizedSet.
func NewRandomizedSet() *RandomizedSet {
	// Инициализируем генератор случайных чисел
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return &RandomizedSet{
		values:  []int{},
		indices: make(map[int]int),
	}
}

// Insert добавляет значение в набор, если его там нет. Возвращает true, если значение добавлено.
func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.indices[val]; exists {
		return false
	}
	rs.values = append(rs.values, val)
	rs.indices[val] = len(rs.values) - 1
	return true
}

// Remove удаляет значение из набора, если оно существует. Возвращает true, если значение удалено.
func (rs *RandomizedSet) Remove(val int) bool {
	idx, exists := rs.indices[val]
	if !exists {
		return false
	}
	lastIdx := len(rs.values) - 1
	lastVal := rs.values[lastIdx]

	// Перемещаем последний элемент на позицию удаляемого
	rs.values[idx] = lastVal
	rs.indices[lastVal] = idx

	// Удаляем последний элемент
	rs.values = rs.values[:lastIdx]
	delete(rs.indices, val)
	return true
}

// GetRandom возвращает случайный элемент из набора.
func (rs *RandomizedSet) GetRandom() int {
	randomIdx := rand.Intn(len(rs.values))
	return rs.values[randomIdx]
}

func main() {
	// Пример использования:
	randomizedSet := NewRandomizedSet()
	fmt.Println(randomizedSet.Insert(1)) // true, 1 вставлено
	fmt.Println(randomizedSet.Remove(2)) // false, 2 отсутствует
	fmt.Println(randomizedSet.Insert(2)) // true, 2 вставлено
	// В наборе теперь [1,2], getRandom может вернуть 1 или 2
	fmt.Println(randomizedSet.GetRandom())
	fmt.Println(randomizedSet.Remove(1)) // true, 1 удалено
	fmt.Println(randomizedSet.Insert(2)) // false, 2 уже присутствует
	// Теперь набор содержит [2], getRandom всегда вернет 2
	fmt.Println(randomizedSet.GetRandom())
}
