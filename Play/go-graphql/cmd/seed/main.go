package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/neogan74/dev-experiments/Play/go-graphql/graph/model"
	"github.com/neogan74/dev-experiments/Play/go-graphql/internal/data"
)

func main() {
	outPath := flag.String("out", "data/seed.json", "output path for the generated seed JSON")
	usersCount := flag.Int("users", 5, "number of users to generate")
	todosPerUser := flag.Int("todos", 4, "number of todos per user")
	randomSeed := flag.Int64("seed", 0, "optional seed for deterministic output (default uses current time)")
	flag.Parse()

	seed := *randomSeed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	gofakeit.Seed(seed)

	if *usersCount <= 0 {
		log.Fatal("users must be greater than zero")
	}
	if *todosPerUser < 0 {
		log.Fatal("todos must be zero or greater")
	}

	users := make([]*model.User, 0, *usersCount)
	for i := 0; i < *usersCount; i++ {
		users = append(users, &model.User{
			ID:   uuid.NewString(),
			Name: gofakeit.Name(),
		})
	}

	todos := make([]data.SeedTodo, 0, (*usersCount)*(*todosPerUser))
	for _, user := range users {
		for i := 0; i < *todosPerUser; i++ {
			todos = append(todos, data.SeedTodo{
				ID:     uuid.NewString(),
				Text:   gofakeit.Sentence(),
				Done:   gofakeit.Bool(),
				UserID: user.ID,
			})
		}
	}

	seedData := data.SeedData{
		Users: users,
		Todos: todos,
	}

	if err := os.MkdirAll(filepath.Dir(*outPath), 0o755); err != nil {
		log.Fatalf("create directory for %s: %v", *outPath, err)
	}

	file, err := os.Create(*outPath)
	if err != nil {
		log.Fatalf("create seed file: %v", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("close file: %v", cerr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(seedData); err != nil {
		log.Fatalf("encode seed data: %v", err)
	}

	fmt.Printf("Generated %d users and %d todos to %s (seed %d)\n", len(users), len(todos), *outPath, seed)
}
