package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/neogan74/dev-experiments/Play/go-graphql/graph/model"
)

// SeedTodo represents the Todo entry in the seed file.
type SeedTodo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
}

// SeedData contains all users and todos that can be loaded into the store.
type SeedData struct {
	Users []*model.User `json:"users"`
	Todos []SeedTodo    `json:"todos"`
}

// Store keeps the application data in-memory in a thread-safe way.
type Store struct {
	mu        sync.RWMutex
	users     map[string]*model.User
	todos     map[string]*model.Todo
	todoOrder []string
}

// NewStore creates an empty store.
func NewStore() *Store {
	return &Store{
		users:     make(map[string]*model.User),
		todos:     make(map[string]*model.Todo),
		todoOrder: make([]string, 0),
	}
}

// LoadFromFile populates the store using the JSON seed file at the provided path.
// Missing files are treated as an empty dataset.
func (s *Store) LoadFromFile(path string) error {
	if path == "" {
		return errors.New("seed file path cannot be empty")
	}

	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("open seed file: %w", err)
	}
	defer func() { _ = f.Close() }()

	return s.LoadFromReader(f)
}

// LoadFromReader populates the store using JSON data read from the reader.
func (s *Store) LoadFromReader(r io.Reader) error {
	var seed SeedData
	if err := json.NewDecoder(r).Decode(&seed); err != nil {
		return fmt.Errorf("decode seed data: %w", err)
	}

	return s.LoadSeedData(seed)
}

// LoadSeedData replaces all existing data in the store using the provided seed.
func (s *Store) LoadSeedData(seed SeedData) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := make(map[string]*model.User, len(seed.Users))
	for _, u := range seed.Users {
		if u == nil {
			continue
		}
		// Create a copy so callers cannot mutate internal state.
		userCopy := *u
		users[u.ID] = &userCopy
	}

	todos := make(map[string]*model.Todo, len(seed.Todos))
	order := make([]string, 0, len(seed.Todos))
	for _, td := range seed.Todos {
		user, ok := users[td.UserID]
		if !ok {
			return fmt.Errorf("todo %q references unknown user %q", td.ID, td.UserID)
		}

		todoCopy := &model.Todo{
			ID:   td.ID,
			Text: td.Text,
			Done: td.Done,
			User: user,
		}
		todos[todoCopy.ID] = todoCopy
		order = append(order, todoCopy.ID)
	}

	s.users = users
	s.todos = todos
	s.todoOrder = order
	return nil
}

// Todos returns all todos in insertion order.
func (s *Store) Todos() []*model.Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.Todo, 0, len(s.todoOrder))
	for _, id := range s.todoOrder {
		if todo, ok := s.todos[id]; ok {
			// Return a shallow copy to keep internal state protected.
			copyTodo := *todo
			result = append(result, &copyTodo)
		}
	}
	return result
}

// Users returns all users currently in the store.
func (s *Store) Users() []*model.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.User, 0, len(s.users))
	for _, user := range s.users {
		userCopy := *user
		result = append(result, &userCopy)
	}
	return result
}

// CreateTodo stores a new todo for an existing user.
func (s *Store) CreateTodo(input model.NewTodo) (*model.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[input.UserID]
	if !ok {
		return nil, fmt.Errorf("unknown user %q", input.UserID)
	}

	todo := &model.Todo{
		ID:   uuid.NewString(),
		Text: input.Text,
		Done: false,
		User: user,
	}

	s.todos[todo.ID] = todo
	s.todoOrder = append(s.todoOrder, todo.ID)
	todoCopy := *todo
	return &todoCopy, nil
}
