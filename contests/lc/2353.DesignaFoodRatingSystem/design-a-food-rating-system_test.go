package designafoodratingsystem

import (
	"testing"
)

func TestFoodRatings(t *testing.T) {
	foods := []string{"sushi", "steak", "pizza", "burger"}
	cuisines := []string{"japanese", "american", "italian", "american"}
	ratings := []int{15, 18, 12, 14}

	obj := Constructor(foods, cuisines, ratings)

	// Проверяем начальные значения
	if got := obj.HighestRated("american"); got != "steak" {
		t.Errorf("HighestRated() = %v, want %v", got, "steak")
	}

	if got := obj.HighestRated("japanese"); got != "sushi" {
		t.Errorf("HighestRated() = %v, want %v", got, "sushi")
	}

	// Изменяем рейтинг
	obj.ChangeRating("burger", 20)

	// Проверяем обновленные значения
	if got := obj.HighestRated("american"); got != "burger" {
		t.Errorf("HighestRated() = %v, want %v", got, "burger")
	}

	// Восстанавливаем предыдущий рейтинг для проверки обратного изменения
	obj.ChangeRating("burger", 14)

	if got := obj.HighestRated("american"); got != "steak" {
		t.Errorf("HighestRated() = %v, want %v", got, "steak")
	}

	// Изменяем рейтинг самого лучшего ресторана
	obj.ChangeRating("steak", 10)

	if got := obj.HighestRated("american"); got != "burger" {
		t.Errorf("HighestRated() = %v, want %v", got, "burger")
	}

	// Проверяем, что рейтинг обновился в глобальной таблице
	if got := obj.g["steak"].rating; got != 10 {
		t.Errorf("Rating for steak = %v, want %v", got, 10)
	}
}
