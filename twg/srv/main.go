package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct{}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	handler := &PlayerServer(&InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":5000", handler))
}
