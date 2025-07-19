package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()

	// Тест поиска слова в пустом Trie
	if trie.Search("apple") {
		t.Errorf("Ожидается false, получено true")
	}

	// Тест вставки и поиска
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("Ожидается true, получено false")
	}

	// Тест поиска префикса
	if !trie.StartsWith("app") {
		t.Errorf("Ожидается true, получено false")
	}

	// Тест поиска слова, которое является префиксом другого слова
	if trie.Search("app") {
		t.Errorf("Ожидается false, получено true")
	}

	// Тест вставки и поиска другого слова
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("Ожидается true, получено false")
	}

	// Тест поиска слова после вставки нескольких слов
	if !trie.StartsWith("a") {
		t.Errorf("Ожидается true, получено false")
	}

	// Тест поиска слова, которое не существует
	if trie.Search("application") {
		t.Errorf("Ожидается false, получено true")
	}

	// Тест длинного префикса
	trie.Insert("application")
	if !trie.Search("application") {
		t.Errorf("Ожидается true, получено false")
	}

	if !trie.StartsWith("appl") {
		t.Errorf("Ожидается true, получено false")
	}
}
